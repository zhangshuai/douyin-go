package douyingo

import (
	"context"
	"net/url"

	"github.com/zhangshuai/douyin-go/conf"
)

// DataExternalItemBaseReq 视频基础数据请求
type DataExternalItemBaseReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
}

// DataExternalItemBase 视频基础数据
type DataExternalItemBase struct {
	AvgPlayDuration float64 `json:"avg_play_duration"` // 30天平均播放时长
	TotalComment    int64   `json:"total_comment"`     // 30天评论数
	TotalLike       int64   `json:"total_like"`        // 30天点赞数
	TotalPlay       int64   `json:"total_play"`        // 30天播放次数
	TotalShare      int64   `json:"total_share"`       // 30天分享数
}

// DDataExternalItemBaseData 视频基础数据
type DDataExternalItemBaseData struct {
	Result DataExternalItemBase `json:"result"` // 视频基础数据
	DYError
}

// DataExternalItemBaseRes 视频基础数据
type DataExternalItemBaseRes struct {
	Data  DDataExternalItemBaseData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// DataExternalItemBase 获取视频基础数据
func (m *Manager) DataExternalItemBase(req DataExternalItemBaseReq) (res DataExternalItemBaseRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s", conf.API_DATA_EXTERNAL_ITEM_BASE, req.AccessToken, req.OpenId, itemId), nil, nil)
	return res, err
}

// DataExternalItemLikeReq 视频点赞数据请求
type DataExternalItemLikeReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

// DataExternalItemLike 视频点赞数据
type DataExternalItemLike struct {
	Date string `json:"date"`
	Like int64  `json:"like"`
}

// DataExternalItemLikeData 视频点赞数据
type DataExternalItemLikeData struct {
	ResultList []DataExternalItemLike `json:"result_list"` // 点赞数据列表
	DYError
}

// DataExternalItemLikeRes 视频点赞
type DataExternalItemLikeRes struct {
	Data  DataExternalItemLikeData `json:"data"`  // 日期
	Extra DYExtra                  `json:"extra"` // 每日点赞数
}

// DataExternalItemLike 获取视频点赞数据
func (m *Manager) DataExternalItemLike(req DataExternalItemLikeReq) (res DataExternalItemLikeRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_LIKE, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}

// DataExternalItemCommentReq 视频评论数据
type DataExternalItemCommentReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

// DataExternalItemComment 视频评论数据
type DataExternalItemComment struct {
	Date    string `json:"date"`    // 日期
	Comment int64  `json:"comment"` // 每日评论数
}

// DataExternalItemCommentData 视频评论数据
type DataExternalItemCommentData struct {
	ResultList []DataExternalItemComment `json:"result_list"` // 评论数据列表
	DYError
}

// DataExternalItemCommentRes 视频评论数据
type DataExternalItemCommentRes struct {
	Data  DataExternalItemCommentData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// DataExternalItemComment 获取视频评论数据
func (m *Manager) DataExternalItemComment(req DataExternalItemCommentReq) (res DataExternalItemCommentRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_COMMENT, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}

// DataExternalItemPlayReq 视频播放数据
type DataExternalItemPlayReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

// DataExternalItemPlay 视频播放数据
type DataExternalItemPlay struct {
	Date string `json:"date"` // 日期
	Play int64  `json:"play"` // 每日播放数
}

// DataExternalItemPlayData 视频播放数据
type DataExternalItemPlayData struct {
	ResultList []DataExternalItemPlay `json:"result_list"` // 播放数据列表
	DYError
}

// DataExternalItemPlayRes 视频播放数据
type DataExternalItemPlayRes struct {
	Data  DataExternalItemPlayData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// DataExternalItemPlay 获取视频播放数据
func (m *Manager) DataExternalItemPlay(req DataExternalItemPlayReq) (res DataExternalItemPlayRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_PLAY, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}

// DataExternalItemShareReq 视频分享数据请求
type DataExternalItemShareReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

// DataExternalItemShare 视频分享数据
type DataExternalItemShare struct {
	Date  string `json:"date"`  // 日期
	Share int64  `json:"share"` // 每日分享数
}

// DataExternalItemShareData 视频分享数据
type DataExternalItemShareData struct {
	ResultList []DataExternalItemShare `json:"result_list"` // 分享数据列表
	DYError
}

// DataExternalItemShareRes 视频分享数据
type DataExternalItemShareRes struct {
	Data  DataExternalItemShareData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// DataExternalItemShare 获取视频分享数据
func (m *Manager) DataExternalItemShare(req DataExternalItemShareReq) (res DataExternalItemShareRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_SHARE, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}
