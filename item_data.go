package douyinGo

import (
	"context"
	"net/url"

	"github.com/zhangshuai/douyin-go/conf"
)

type DataExternalItemBaseReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
}

type DataExternalItemBase struct {
	AvgPlayDuration float64 `json:"avg_play_duration"` // 30天平均播放时长
	TotalComment    int64   `json:"total_comment"`     // 30天评论数
	TotalLike       int64   `json:"total_like"`        // 30天点赞数
	TotalPlay       int64   `json:"total_play"`        // 30天播放次数
	TotalShare      int64   `json:"total_share"`       // 30天分享数
}

type DDataExternalItemBaseData struct {
	Result DataExternalItemBase `json:"result"` // 视频基础数据
	DYError
}

type DataExternalItemBaseRes struct {
	Data  DDataExternalItemBaseData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// 获取视频基础数据
func (m *Manager) DataExternalItemBase(req DataExternalItemBaseReq) (res DataExternalItemBaseRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s", conf.API_DATA_EXTERNAL_ITEM_BASE, req.AccessToken, req.OpenId, itemId), nil, nil)
	return res, err
}

type DataExternalItemLikeReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

type DataExternalItemLike struct {
	Date string `json:"date"`
	Like int64  `json:"like"`
}

type DataExternalItemLikeData struct {
	ResultList []DataExternalItemLike `json:"result_list"` // 点赞数据列表
	DYError
}

type DataExternalItemLikeRes struct {
	Data  DataExternalItemLikeData `json:"data"`  // 日期
	Extra DYExtra                  `json:"extra"` // 每日点赞数
}

// 获取视频点赞数据
func (m *Manager) DataExternalItemLike(req DataExternalItemLikeReq) (res DataExternalItemLikeRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_LIKE, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}

type DataExternalItemCommentReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

type DataExternalItemComment struct {
	Date    string `json:"date"`    // 日期
	Comment int64  `json:"comment"` // 每日评论数
}

type DataExternalItemCommentData struct {
	ResultList []DataExternalItemComment `json:"result_list"` // 评论数据列表
	DYError
}

type DataExternalItemCommentRes struct {
	Data  DataExternalItemCommentData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// 获取视频评论数据
func (m *Manager) DataExternalItemComment(req DataExternalItemCommentReq) (res DataExternalItemCommentRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_COMMENT, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}

type DataExternalItemPlayReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

type DataExternalItemPlay struct {
	Date string `json:"date"` // 日期
	Play int64  `json:"play"` // 每日播放数
}

type DataExternalItemPlayData struct {
	ResultList []DataExternalItemPlay `json:"result_list"` // 播放数据列表
	DYError
}

type DataExternalItemPlayRes struct {
	Data  DataExternalItemPlayData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// 获取视频播放数据
func (m *Manager) DataExternalItemPlay(req DataExternalItemPlayReq) (res DataExternalItemPlayRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_PLAY, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}

type DataExternalItemShareReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ItemId      string // item_id，仅能查询access_token对应用户上传的视频
	DateType    int64  // 近7/15天；输入7代表7天、15代表15天
}

type DataExternalItemShare struct {
	Date  string `json:"date"`  // 日期
	Share int64  `json:"share"` // 每日分享数
}

type DataExternalItemShareData struct {
	ResultList []DataExternalItemShare `json:"result_list"` // 分享数据列表
	DYError
}

type DataExternalItemShareRes struct {
	Data  DataExternalItemShareData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// 获取视频分享数据
func (m *Manager) DataExternalItemShare(req DataExternalItemShareReq) (res DataExternalItemShareRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&item_id=%s&date_type=%d", conf.API_DATA_EXTERNAL_ITEM_SHARE, req.AccessToken, req.OpenId, itemId, req.DateType), nil, nil)
	return res, err
}
