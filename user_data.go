package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// DataExternalUserItemReq 用户视频情况请求
type DataExternalUserItemReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

// DataExternalUserItem 用户视频情况
type DataExternalUserItem struct {
	Date       string `json:"date"`        // 日期
	NewIssue   int64  `json:"new_issue"`   // 每日发布内容数
	NewPlay    int64  `json:"new_play"`    // 每天新增视频播放
	TotalIssue int64  `json:"total_issue"` // 每日内容总数
}

// DataExternalUserItemData 用户视频情况
type DataExternalUserItemData struct {
	ResultList []DataExternalUserItem `json:"result_list"` // 用户视频数据
	DYError
}

// DataExternalUserItemRes 用户视频情况
type DataExternalUserItemRes struct {
	Data  DataExternalUserItemData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// DataExternalUserItem 获取用户视频情况
func (m *Manager) DataExternalUserItem(req DataExternalUserItemReq) (res DataExternalUserItemRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_ITEM, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

// DataExternalUserFansReq 用户粉丝数请求
type DataExternalUserFansReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

// DataExternalUserFans 用户粉丝数
type DataExternalUserFans struct {
	Date      string `json:"date"`       // 日期
	NewFans   int64  `json:"new_fans"`   // 每天新粉丝数
	TotalFans int64  `json:"total_fans"` // 每日总粉丝数
}

// DataExternalUserFansData 用户粉丝数
type DataExternalUserFansData struct {
	ResultList []DataExternalUserFans `json:"result_list"` // 用户粉丝数据
	DYError
}

// DataExternalUserFansRes 用户粉丝数
type DataExternalUserFansRes struct {
	Data  DataExternalUserFansData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// DataExternalUserFans 获取用户粉丝数
func (m *Manager) DataExternalUserFans(req DataExternalUserFansReq) (res DataExternalUserFansRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_FANS, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

// DataExternalUserLikeReq 用户点赞数请求
type DataExternalUserLikeReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

// DataExternalUserLike 用户点赞数
type DataExternalUserLike struct {
	Date    string `json:"date"`     // 日期
	NewLike int64  `json:"new_like"` // 新增点赞
}

// DataExternalUserLikeData 用户点赞数
type DataExternalUserLikeData struct {
	ResultList []DataExternalUserLike `json:"result_list"` // 用户点赞数据
	DYError
}

// DataExternalUserLikeRes 用户点赞数
type DataExternalUserLikeRes struct {
	Data  DataExternalUserLikeData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// DataExternalUserLike 获取用户点赞数
func (m *Manager) DataExternalUserLike(req DataExternalUserLikeReq) (res DataExternalUserLikeRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_LIKE, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

// DataExternalUserCommentReq 用户评论数请求
type DataExternalUserCommentReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

// DataExternalUserComment 用户评论数
type DataExternalUserComment struct {
	Date       string `json:"date"`
	NewComment int64  `json:"new_comment"`
}

// DataExternalUserCommentData 用户评论数
type DataExternalUserCommentData struct {
	ResultList []DataExternalUserComment `json:"result_list"` // 用户评论数据
	DYError
}

// DataExternalUserCommentRes 用户评论数
type DataExternalUserCommentRes struct {
	Data  DataExternalUserCommentData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// DataExternalUserComment 获取用户评论数
func (m *Manager) DataExternalUserComment(req DataExternalUserCommentReq) (res DataExternalUserCommentRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_COMMENT, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

// DataExternalUserShareReq 用户分享数请求
type DataExternalUserShareReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

// DataExternalUserShare 用户分享数
type DataExternalUserShare struct {
	Date     string `json:"date"`      // 日期
	NewShare int64  `json:"new_share"` // 新增分享
}

// DataExternalUserShareData 用户分享数
type DataExternalUserShareData struct {
	ResultList []DataExternalUserShare `json:"result_list"` // 用户分享数据
	DYError
}

// DataExternalUserShareRes 用户分享数
type DataExternalUserShareRes struct {
	Data  DataExternalUserShareData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// DataExternalUserShare 获取用户分享数
func (m *Manager) DataExternalUserShare(req DataExternalUserShareReq) (res DataExternalUserShareRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_SHARE, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}

// DataExternalUserProfileReq 用户主页访问数请求
type DataExternalUserProfileReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	DataType    int64  // 近7/15天；输入7代表7天、15代表15天、30代表30天
}

// DataExternalUserProfile 用户主页访问数
type DataExternalUserProfile struct {
	Date      string `json:"date"`       // 日期
	ProfileUV int64  `json:"profile_uv"` // 当日个人主页访问人数
}

// DataExternalUserProfileData 用户主页访问数
type DataExternalUserProfileData struct {
	ResultList []DataExternalUserProfile `json:"result_list"` // 用户主页访问数据
	DYError
}

// DataExternalUserProfileRes 用户主页访问数
type DataExternalUserProfileRes struct {
	Data  DataExternalUserProfileData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// DataExternalUserProfile 获取用户主页访问数
func (m *Manager) DataExternalUserProfile(req DataExternalUserProfileReq) (res DataExternalUserProfileRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_USER_PROFILE, req.AccessToken, req.OpenId, req.DataType), nil, nil)
	return res, err
}
