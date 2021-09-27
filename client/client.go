package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"strings"

	"github.com/zhangshuai/douyin-go/internal/log"
	"github.com/zhangshuai/douyin-go/reqid"
)

// UserAgent UA
var UserAgent = "Golang douyin/client package"

// DefaultClient 默认Client
var DefaultClient = Client{&http.Client{Transport: http.DefaultTransport}}

// DebugMode 用来打印调试信息
var DebugMode = false

// DeepDebugInfo 调试信息
var DeepDebugInfo = false

// --------------------------------------------------------------------

// Client 负责发送HTTP请求到抖音接口服务器
type Client struct {
	*http.Client
}

// TurnOnDebug 开启Debug模式
func TurnOnDebug() {
	DebugMode = true
}

func newRequest(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, reqUrl, body)
	if err != nil {
		return
	}

	if headers == nil {
		headers = http.Header{}
	}

	req.Header = headers
	req = req.WithContext(ctx)

	if DebugMode {
		trace := &httptrace.ClientTrace{
			GotConn: func(connInfo httptrace.GotConnInfo) {
				remoteAddr := connInfo.Conn.RemoteAddr()
				log.Debug(fmt.Sprintf("Network: %s, Remote ip:%s, URL: %s", remoteAddr.Network(), remoteAddr.String(), req.URL))
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		bs, bErr := httputil.DumpRequest(req, DeepDebugInfo)
		if bErr != nil {
			err = bErr
			return
		}
		log.Debug(string(bs))
	}
	return
}

// DoRequestWith 请求
func (r Client) DoRequestWith(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqUrl, headers, body)
	if err != nil {
		return
	}
	req.ContentLength = int64(bodyLength)
	return r.Do(ctx, req)
}

// DoRequestWithJson JSON请求
func (r Client) DoRequestWithJson(ctx context.Context, method, reqUrl string, headers http.Header,
	data interface{}) (resp *http.Response, err error) {

	var reqBody []byte
	if data != nil {
		reqBody, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("Content-Type", "application/json")
	return r.DoRequestWith(ctx, method, reqUrl, headers, bytes.NewReader(reqBody), len(reqBody))
}

// Do 请求
func (r Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if reqId, ok := reqid.FromContext(ctx); ok {
		req.Header.Set("X-Reqid", reqId)
	}

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", UserAgent)
	}

	transport := r.Transport // don't change r.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	// avoid cancel() is called before Do(req), but isn't accurate
	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	default:
	}

	if tr, ok := getRequestCanceler(transport); ok {
		// support CancelRequest
		reqC := make(chan bool, 1)
		go func() {
			resp, err = r.Client.Do(req)
			reqC <- true
		}()
		select {
		case <-reqC:
		case <-ctx.Done():
			tr.CancelRequest(req)
			<-reqC
			err = ctx.Err()
		}
	} else {
		resp, err = r.Client.Do(req)
	}
	return
}

// ErrorInfo 错误信息
type ErrorInfo struct {
	Err   string `json:"error,omitempty"`
	Key   string `json:"key,omitempty"`
	Reqid string `json:"reqid,omitempty"`
	Errno int    `json:"errno,omitempty"`
	Code  int    `json:"code"`
}

// Error 错误
func (r *ErrorInfo) Error() string {
	return r.Err
}

func parseError(e *ErrorInfo, r io.Reader) {
	body, err1 := ioutil.ReadAll(r)
	if err1 != nil {
		e.Err = err1.Error()
		return
	}

	var ret struct {
		Err   string `json:"error"`
		Key   string `json:"key"`
		Errno int    `json:"errno"`
	}
	if json.Unmarshal(body, &ret) == nil && ret.Err != "" {
		e.Err, e.Key, e.Errno = ret.Err, ret.Key, ret.Errno
		return
	}
	e.Err = string(body)
}

// ResponseError 错误响应
func ResponseError(resp *http.Response) error {
	e := &ErrorInfo{
		Reqid: resp.Header.Get("X-Reqid"),
		Code:  resp.StatusCode,
	}
	if resp.StatusCode > 299 {
		if resp.ContentLength != 0 {
			ct, ok := resp.Header["Content-Type"]
			if ok && strings.HasPrefix(ct[0], "application/json") {
				parseError(e, resp.Body)
			} else {
				bs, rErr := ioutil.ReadAll(resp.Body)
				if rErr != nil {
					e.Err = rErr.Error()
				}
				e.Err = strings.TrimRight(string(bs), "\n")
			}
		} else if resp.Status != "" {
			return errors.New(resp.Status)
		}
	}
	return e
}

// CallRet Http请求
func CallRet(ctx context.Context, ret interface{}, resp *http.Response) (err error) {
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if DebugMode {
		bs, dErr := httputil.DumpResponse(resp, DeepDebugInfo)
		if dErr != nil {
			err = dErr
			return
		}
		log.Debug(string(bs))
	}
	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength != 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			if err != nil {
				return
			}
		}
		if resp.StatusCode == 200 {
			return nil
		}
	}
	return ResponseError(resp)
}

// CallWithJson JSON请求
func (r Client) CallWithJson(ctx context.Context, ret interface{}, method, reqUrl string, headers http.Header,
	param interface{}) (err error) {

	resp, err := r.DoRequestWithJson(ctx, method, reqUrl, headers, param)
	if err != nil {
		return err
	}
	return CallRet(ctx, ret, resp)
}

// DoRequestWith64 请求
func (r Client) DoRequestWith64(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int64) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqUrl, headers, body)
	if err != nil {
		return
	}
	req.ContentLength = bodyLength
	return r.Do(ctx, req)
}

// CallWith64 请求
func (r Client) CallWith64(ctx context.Context, ret interface{}, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int64) (err error) {

	resp, err := r.DoRequestWith64(ctx, method, reqUrl, headers, body, bodyLength)
	if err != nil {
		return err
	}
	return CallRet(ctx, ret, resp)
}

type requestCanceler interface {
	CancelRequest(req *http.Request)
}

type nestedObjectGetter interface {
	NestedObject() interface{}
}

func getRequestCanceler(tp http.RoundTripper) (rc requestCanceler, ok bool) {
	if rc, ok = tp.(requestCanceler); ok {
		return
	}

	p := interface{}(tp)
	for {
		getter, ok1 := p.(nestedObjectGetter)
		if !ok1 {
			return
		}
		p = getter.NestedObject()
		if rc, ok = p.(requestCanceler); ok {
			return
		}
	}
}
