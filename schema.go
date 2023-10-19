package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// SchemaGetShareReq H5分享跳转链接获取请求
type SchemaGetShareReq struct {
	AccessToken string             // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	Body        SchemaGetShareBody // H5分享跳转链接获取body
}

// SchemaGetShareBody H5分享跳转链接获取body
type SchemaGetShareBody struct {
	ClientTicket   string    `json:"client_ticket"`              // openTicket获取的ticket
	ExpireAt       int64     `json:"expire_at"`                  // 过期时间
	HashtagList    []string  `json:"hashtag_list,omitempty"`     // 支持有第三方预设内容分享抖音时默认携带的话题
	ImageListPath  []string  `json:"image_list_path,omitempty"`  // 图片文件路径(多个)，图集模式分享
	ImagePath      string    `json:"image_path,omitempty"`       // 图片文件路径
	MicroAppInfo   *MicroApp `json:"micro_app_info,omitempty"`   // 添加小程序。视频成功发布视频后，在视频左下角带有小程序入口。
	PoiId          string    `json:"poi_id,omitempty"`           // 地理位置信息锚点 id，与小程序 appId 互斥，优先展示小程序。
	ShareToPublish string    `json:"share_to_publish,omitempty"` // 为1时直接分享到抖音发布页（仅视频）
	State          string    `json:"state,omitempty"`            // 建议填写，按照文档获取share_id，可以获取视频发布情况
	Title          string    `json:"title,omitempty"`            // 视频标题
	VideoPath      string    `json:"video_path,omitempty"`       // 视频文件路径（单个，不能超过128M)。
}

// SchemaGetShareData H5分享跳转链接获取
type SchemaGetShareData struct {
	Schema string `json:"schema"` // Schema链接
}

// SchemaGetShareRes H5分享跳转链接获取
type SchemaGetShareRes struct {
	Data SchemaGetShareData `json:"data"`
	DYExtraV1
}

// SchemaGetShare H5分享跳转链接获取
func (m *Manager) SchemaGetShare(req SchemaGetShareReq) (res SchemaGetShareRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_SCHEMA_GET_SHARE, req.AccessToken), nil, req.Body)
	return res, err
}

// SchemaGetUserProfileReq 个人页跳转链接获取
type SchemaGetUserProfileReq struct {
	AccessToken string                   // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	Body        SchemaGetUserProfileBody // 个人页跳转链接获取body
}

// SchemaGetUserProfileBody 个人页跳转链接获取body
type SchemaGetUserProfileBody struct {
	Account  string `json:"account,omitempty"` // 抖音号
	ExpireAt int64  `json:"expire_at"`         // 生成短链过期时间
	OpenId   string `json:"open_id,omitempty"` // open id
}

// SchemaGetUserProfileData 个人页跳转链接获取
type SchemaGetUserProfileData struct {
	Schema string `json:"schema"` // Schema链接
}

// SchemaGetUserProfileRes 个人页跳转链接获取
type SchemaGetUserProfileRes struct {
	Data SchemaGetUserProfileData `json:"data"`
	DYExtraV1
}

// SchemaGetUserProfile 个人页跳转链接获取
func (m *Manager) SchemaGetUserProfile(req SchemaGetUserProfileReq) (res SchemaGetUserProfileRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_SCHEMA_GET_USER_PROFILE, req.AccessToken), nil, req.Body)
	return res, err
}

// SchemaGetChatReq 个人会话页跳转链接获取
type SchemaGetChatReq struct {
	AccessToken string            // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	Body        SchemaGetChatBody // 个人会话页跳转链接获取body
}

// SchemaGetChatBody 个人会话页跳转链接获取body
type SchemaGetChatBody struct {
	Account  string `json:"account,omitempty"` // 抖音号
	ExpireAt int64  `json:"expire_at"`         // 生成短链过期时间
	OpenId   string `json:"open_id,omitempty"` // open id
}

// SchemaGetChatData 个人会话页跳转链接获取
type SchemaGetChatData struct {
	Schema string `json:"schema"` // Schema链接
}

// SchemaGetChatRes 个人会话页跳转链接获取
type SchemaGetChatRes struct {
	Data SchemaGetChatData `json:"data"`
	DYExtraV1
}

// SchemaGetChat 个人会话页跳转链接获取
func (m *Manager) SchemaGetChat(req SchemaGetChatReq) (res SchemaGetChatRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_SCHEMA_GET_CHAT, req.AccessToken), nil, req.Body)
	return res, err
}

// SchemaGetItemInfoReq 视频详情页跳转链接获取
type SchemaGetItemInfoReq struct {
	AccessToken string                // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	Body        SchemaGetItemInfoBody // 视频详情页跳转链接获取body
}

// SchemaGetItemInfoBody 视频详情页跳转链接获取body
type SchemaGetItemInfoBody struct {
	ExpireAt int64  `json:"expire_at"`          // 生成短链过期时间
	ItemId   string `json:"item_id,omitempty"`  // 视频id
	VideoId  string `json:"video_id,omitempty"` // 视频id
}

// SchemaGetItemInfoData 视频详情页跳转链接获取
type SchemaGetItemInfoData struct {
	Schema string `json:"schema"` // Schema链接
}

// SchemaGetItemInfoRes 视频详情页跳转链接获取
type SchemaGetItemInfoRes struct {
	Data SchemaGetItemInfoData `json:"data"`
	DYExtraV1
}

// SchemaGetItemInfo 视频详情页跳转链接获取
func (m *Manager) SchemaGetItemInfo(req SchemaGetItemInfoReq) (res SchemaGetItemInfoRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_SCHEMA_GET_ITEM_INFO, req.AccessToken), nil, req.Body)
	return res, err
}

// SchemaGetLiveReq 直播间跳转链接获取
type SchemaGetLiveReq struct {
	AccessToken string            // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	Body        SchemaGetLiveBody // 直播间跳转链接获取body
}

// SchemaGetLiveBody 直播间跳转链接获取body
type SchemaGetLiveBody struct {
	Account  string `json:"account,omitempty"` // 抖音号
	ExpireAt int64  `json:"expire_at"`         // 生成短链过期时间
	OpenId   string `json:"open_id,omitempty"` // open id
}

// SchemaGetLiveData 直播间跳转链接获取
type SchemaGetLiveData struct {
	Schema string `json:"schema"` // Schema链接
}

// SchemaGetLiveRes 直播间跳转链接获取
type SchemaGetLiveRes struct {
	Data SchemaGetLiveData `json:"data"`
	DYExtraV1
}

// SchemaGetLive 直播间跳转链接获取
func (m *Manager) SchemaGetLive(req SchemaGetLiveReq) (res SchemaGetLiveRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_SCHEMA_GET_LIVE, req.AccessToken), nil, req.Body)
	return res, err
}
