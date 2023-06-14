package douyingo

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"

	"encoding/json"

	"github.com/zhangshuai/douyin-go/conf"
)

// OpenTicketReq open_ticket请求
type OpenTicketReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// OpenTicketData open_ticket
type OpenTicketData struct {
	ExpiresIn int64  `json:"expires_in"` // open_ticket超时时间，单位秒
	Ticket    string `json:"ticket"`     // open_ticket接口调用凭证
	DYError
}

// OpenTicketRes open_ticket
type OpenTicketRes struct {
	Data  OpenTicketData `json:"data"`
	Extra DYExtra        `json:"extra"`
}

// OpenTicket 获取open_ticket
func (m *Manager) OpenTicket(req OpenTicketReq) (res OpenTicketRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_OPEN_TICKET, req.AccessToken), nil, nil)
	return res, err
}

// OpenSignatureReq 验证签名请求
type OpenSignatureReq struct {
	Ticket    string
	Timestamp int64  // 时间戳
	NonceStr  string // 生成签名用的随机字符串
}

// OpenSignature 验证签名
func (m *Manager) OpenSignature(req OpenSignatureReq) string {
	params := fmt.Sprintf("nonce_str=%s&ticket=%s&timestamp=%d", req.NonceStr, req.Ticket, req.Timestamp)
	fmt.Println(params)
	signature := fmt.Sprintf("%x", md5.Sum([]byte(params)))
	return signature
}

// ShareSchemaReq H5分享Schema请求
type ShareSchemaReq struct {
	ShareType      string   `json:"share_type"`
	ClientKey      string   `json:"client_key"`
	NonceStr       string   `json:"nonce_str"`
	Timestamp      string   `json:"timestamp"`
	Signature      string   `json:"signature"`
	State          string   `json:"state"`
	ImagePath      string   `json:"image_path"`
	ImageListPath  []string `json:"image_list_path"`
	VideoPath      string   `json:"video_path"`
	HashtagList    []string `json:"hashtag_list"`
	ShareToPublish string   `json:"share_to_publish"`
	Title          string   `json:"title"`
	PoiId          string   `json:"poi_id"`
}

// ShareSchema H5分享Schema
func (m *Manager) ShareSchema(req ShareSchemaReq) (string, error) {
	var params []string
	if req.ShareType != "" {
		params = append(params, fmt.Sprintf("share_type=%s", req.ShareType))
	}
	if req.ClientKey != "" {
		params = append(params, fmt.Sprintf("client_key=%s", req.ClientKey))
	}
	if req.NonceStr != "" {
		params = append(params, fmt.Sprintf("nonce_str=%s", req.NonceStr))
	}
	if req.Timestamp != "" {
		params = append(params, fmt.Sprintf("timestamp=%s", req.Timestamp))
	}
	if req.Signature != "" {
		params = append(params, fmt.Sprintf("signature=%s", req.Signature))
	}
	if req.State != "" {
		params = append(params, fmt.Sprintf("state=%s", req.State))
	}
	if req.ImagePath != "" {
		params = append(params, fmt.Sprintf("image_path=%s", url.QueryEscape(req.ImagePath)))
	}
	if len(req.ImageListPath) > 0 {
		imageListPathStr, err := json.Marshal(req.ImageListPath)
		if err != nil {
			return "", err
		}
		params = append(params, fmt.Sprintf("image_list_path=%s", url.QueryEscape(string(imageListPathStr))))
	}
	if req.VideoPath != "" {
		params = append(params, fmt.Sprintf("video_path=%s", url.QueryEscape(req.VideoPath)))
	}
	if len(req.HashtagList) > 0 {
		hashtagListStr, err := json.Marshal(req.HashtagList)
		if err != nil {
			return "", err
		}
		params = append(params, fmt.Sprintf("hashtag_list=%s", url.QueryEscape(string(hashtagListStr))))
	}
	if req.ShareToPublish != "" {
		params = append(params, fmt.Sprintf("share_to_publish=%s", req.ShareToPublish))
	}
	if req.Title != "" {
		params = append(params, fmt.Sprintf("title=%s", url.QueryEscape(req.Title)))
	}
	if req.PoiId != "" {
		params = append(params, fmt.Sprintf("poi_id=%s", req.PoiId))
	}

	return fmt.Sprintf("%s?%s", conf.API_SHARE_SCHEMA, strings.Join(params, "&")), nil
}
