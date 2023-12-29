package douyingo

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"

	"encoding/json"

	"github.com/guaidashu/douyin-go/conf"
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

// MicroApp 小程序信息
type MicroApp struct {
	AppId       string `json:"appId"`       // 小程序appId
	AppTitle    string `json:"appTitle"`    // 小程序标题
	AppUrl      string `json:"appUrl"`      // 小程序中生成该页面时写的path地址5
	Description string `json:"description"` // 小程序描述语
}

// ShareSchemaReq H5分享Schema请求
type ShareSchemaReq struct {
	ShareType      string    `json:"share_type"`       // 固定值为 h5
	ClientKey      string    `json:"client_key"`       // 应用 key，在控制台中应用的总览页面获取
	NonceStr       string    `json:"nonce_str"`        // 随机字符串
	Timestamp      string    `json:"timestamp"`        // 时间戳
	Signature      string    `json:"signature"`        // 验签签名
	State          string    `json:"state"`            // 建议填写 按照查询视频发布结果获取 share_id，可以获取视频发布情况
	ImagePath      string    `json:"image_path"`       // 图片文件路径（单个，不能超过 20M） 当 video_path 存在时优先使用 video_path 当前支持的格式包含 png/jpg/gif 关于该参数请阅读表格下面的注意事项。
	ImageListPath  []string  `json:"image_list_path"`  // 图片文件路径（多个），图集模式分享 当 video_path 存在时优先使用 video_path 当前支持的格式包含 png/jpg。
	VideoPath      string    `json:"video_path"`       // 视频文件路径（单个，不能超过 128M)。 当前支持的格式包含 mp4/mov。 关于该参数请阅读表格下面的注意事项。
	HashtagList    []string  `json:"hashtag_list"`     // 支持有第三方预设内容分享抖音时默认携带的话题，指定的话题会展现在发布页面。 用户可自行删除该话题，该话题类型支持商业化话题和普通话题。发布后和抖音原生话题没有差别。
	MicroAppInfo   *MicroApp `json:"micro_app_info"`   // 添加小程序。视频成功发布视频后，在视频左下角带有小程序入口。
	ShareToPublish string    `json:"share_to_publish"` // 为 1 时直接分享到抖音发布页（仅视频）
	Title          string    `json:"title"`            // 视频标题
	PoiId          string    `json:"poi_id"`           // 地理位置信息锚点 id，与小程序 appId 互斥，优先展示小程序。
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
	if req.MicroAppInfo != nil {
		microAppInfoStr, err := json.Marshal(req.MicroAppInfo)
		if err != nil {
			return "", err
		}
		params = append(params, fmt.Sprintf("micro_app_info=%s", url.QueryEscape(string(microAppInfoStr))))
	}
	if req.PoiId != "" {
		params = append(params, fmt.Sprintf("poi_id=%s", req.PoiId))
	}

	return fmt.Sprintf("%s?%s", conf.API_SHARE_SCHEMA, strings.Join(params, "&")), nil
}
