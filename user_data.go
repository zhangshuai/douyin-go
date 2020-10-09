package douyinGo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

type DataExternalUserItemReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

type DataExternalUserItem struct {
	Date       string `json:"date"`        // 日期
	NewIssue   int64  `json:"new_issue"`   // 每日发布内容数
	NewPlay    int64  `json:"new_play"`    // 每天新增视频播放
	TotalIssue int64  `json:"total_issue"` // 每日内容总数
}

type DataExternalUserItemData struct {
	ResultList []DataExternalUserItem `json:"result_list"` // 用户视频数据
	DYError
}

type DataExternalUserItemRes struct {
	Data  DataExternalUserItemData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// 获取用户视频情况
func (m *Manager) DataExternalUserItem(req DataExternalUserItemReq) (res DataExternalUserItemRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_ITEM, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

type DataExternalUserFansReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

type DataExternalUserFans struct {
	Date      string `json:"date"`       // 日期
	NewFans   int64  `json:"new_fans"`   // 每天新粉丝数
	TotalFans int64  `json:"total_fans"` // 每日总粉丝数
}

type DataExternalUserFansData struct {
	ResultList []DataExternalUserFans `json:"result_list"` // 用户粉丝数据
	DYError
}

type DataExternalUserFansRes struct {
	Data  DataExternalUserFansData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// 获取用户粉丝数
func (m *Manager) DataExternalUserFans(req DataExternalUserFansReq) (res DataExternalUserFansRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_FANS, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

type DataExternalUserLikeReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

type DataExternalUserLike struct {
	Date    string `json:"date"`     // 日期
	NewLike int64  `json:"new_like"` // 新增点赞
}

type DataExternalUserLikeData struct {
	ResultList []DataExternalUserLike `json:"result_list"` // 用户点赞数据
	DYError
}

type DataExternalUserLikeRes struct {
	Data  DataExternalUserLikeData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// 获取用户点赞数
func (m *Manager) DataExternalUserLike(req DataExternalUserLikeReq) (res DataExternalUserLikeRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_LIKE, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

type DataExternalUserCommentReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

type DataExternalUserComment struct {
	Date       string `json:"date"`
	NewComment int64  `json:"new_comment"`
}

type DataExternalUserCommentData struct {
	ResultList []DataExternalUserComment `json:"result_list"` // 用户评论数据
	DYError
}

type DataExternalUserCommentRes struct {
	Data  DataExternalUserCommentData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// 获取用户评论数
func (m *Manager) DataExternalUserComment(req DataExternalUserCommentReq) (res DataExternalUserCommentRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_COMMENT, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

type DataExternalUserShareReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

type DataExternalUserShare struct {
	Date     string `json:"date"`      // 日期
	NewShare int64  `json:"new_share"` // 新增分享
}

type DataExternalUserShareData struct {
	ResultList []DataExternalUserShare `json:"result_list"` // 用户分享数据
	DYError
}

type DataExternalUserShareRes struct {
	Data  DataExternalUserShareData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// 获取用户分享数
func (m *Manager) DataExternalUserShare(req DataExternalUserShareReq) (res DataExternalUserShareRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_SHARE, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

type DataExternalUserProfileReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

type DataExternalUserProfile struct {
	Date      string `json:"date"`       // 日期
	ProfileUV int64  `json:"profile_uv"` // 当日个人主页访问人数
}

type DataExternalUserProfileData struct {
	ResultList []DataExternalUserProfile `json:"result_list"` // 用户主页访问数据
	DYError
}

type DataExternalUserProfileRes struct {
	Data  DataExternalUserProfileData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// 获取用户主页访问数
func (m *Manager) DataExternalUserProfile(req DataExternalUserProfileReq) (res DataExternalUserProfileRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_PROFILE, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}
