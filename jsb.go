package douyingo

import (
	"context"
	"crypto/md5"
	"fmt"

	"github.com/zhangshuai/douyin-go/conf"
)

// JsTicketReq jsapi_ticket请求
type JsTicketReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// JsTicketData jsapi_ticket
type JsTicketData struct {
	ExpiresIn int64  `json:"expires_in"` // access_token接口调用凭证超时时间，单位（秒）
	Ticket    string `json:"ticket"`     // js接口调用凭证
	DYError
}

// JsTicketRes jsapi_ticket
type JsTicketRes struct {
	Data  JsTicketData `json:"data"`
	Extra DYExtra      `json:"extra"`
}

// JsTicket 获取jsapi_ticket
func (m *Manager) JsTicket(req JsTicketReq) (res JsTicketRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_JS_TICKET, req.AccessToken), nil, nil)
	return res, err
}

// ConfigSignReq 验证签名请求
type ConfigSignReq struct {
	JsTicket  string
	Timestamp int64  // 时间戳
	NonceStr  string // 生成签名用的随机字符串
	Url       string // 为应用申请的JSB安全域名，需要携带协议。
}

// JsConfigSignature 通过config方法验证签名
func (m *Manager) JsConfigSignature(req ConfigSignReq) string {
	params := fmt.Sprintf("jsapi_ticket=%s&nonce_str=%s&timestamp=%d&url=%s", req.JsTicket, req.NonceStr, req.Timestamp, req.Url)
	signature := fmt.Sprintf("%x", md5.Sum([]byte(params)))
	return signature
}
