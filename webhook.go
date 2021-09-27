package douyingo

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/zhangshuai/douyin-go/conf"
)

// WebhookEvent Webhook事件
type WebhookEvent struct {
	Event      string         `json:"event"`                  // 事件
	ClientKey  string         `json:"client_key"`             // 使用应用的client_key
	FromUserId string         `json:"from_user_id,omitempty"` // 事件发起用户user_id
	ToUserId   string         `json:"to_user_id,omitempty"`   // 事件接收用户user_id
	Content    WebhookContent `json:"content,omitempty"`      // 不同的event对应不同的content
}

// WebhookContent Webhook内容
type WebhookContent struct {
	Challenge   int      `json:"challenge,omitempty"`   // webhook验证码
	ItemId      string   `json:"item_id,omitempty"`     // 视频id
	ShareId     string   `json:"share_id,omitempty"`    // 分享id
	Scopes      []string `json:"scopes,omitempty"`      // 授权scope列表
	Description string   `json:"discription,omitempty"` // 具体见私信事件列表
	Scene       string   `json:"scene,omitempty"`       // 进入对话来源场景["video", "homepage"]
	Object      string   `json:"object,omitempty"`      // 来源场景对应id（video对应视频id）
}

// WebhookSignature Webhook签名
func (m *Manager) WebhookSignature(body []byte, sign string) bool {
	body = append([]byte(m.Credentials.ClientSecret), body...)
	return fmt.Sprintf("%x", sha1.Sum(body)) == sign
}

// EventStatusListReq 事件订阅状态请求
type EventStatusListReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// EventStatus 事件订阅状态
type EventStatus struct {
	Event  string `json:"event"`  // 推送事件名称
	Status int64  `json:"status"` // 事件订阅状态 * `0` - 未订阅 * `1` - 已订阅
}

// EventStatusListResData 事件订阅状态
type EventStatusListResData struct {
	List []EventStatus `json:"list"` // 事件列表
	DYError
}

// EventStatusListRes 事件订阅状态
type EventStatusListRes struct {
	Data  EventStatusListResData `json:"data"`
	Extra DYExtra                `json:"extra"`
}

// EventStatusList 获取事件订阅状态
func (m *Manager) EventStatusList(req EventStatusListReq) (res EventStatusListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_EVENT_STATUS_LIST, req.AccessToken), nil, nil)
	return res, err
}

// EventStatusUpdateReq 事件订阅状态
type EventStatusUpdateReq struct {
	AccessToken string                // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Body        EventStatusUpdateBody // 更新事件订阅请求body
}

// EventStatusUpdateBody 事件订阅状态
type EventStatusUpdateBody struct {
	List []EventStatus `json:"list"` // 更新事件列表
}

// EventStatusUpdateResData 事件订阅状态
type EventStatusUpdateResData struct {
	DYError
}

// EventStatusUpdateRes 事件订阅状态
type EventStatusUpdateRes struct {
	Data  EventStatusUpdateResData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// EventStatusUpdate 更新应用推送事件订阅状态
func (m *Manager) EventStatusUpdate(req EventStatusUpdateReq) (res EventStatusUpdateRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_EVENT_STATUS_UPDATE, req.AccessToken), nil, req.Body)
	return res, err
}
