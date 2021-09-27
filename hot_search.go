package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// HotSearchSentencesReq 实时热点词请求
type HotSearchSentencesReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// HotSearchSentence 实时热点词
type HotSearchSentence struct {
	HotLevel int64  `json:"hot_level"` // 热度 综合点赞、评论、转发等计算得出
	Sentence string `json:"sentence"`  // 热点词
}

// HotSearchSentencesData 实时热点词
type HotSearchSentencesData struct {
	ActiveTime string              `json:"active_time"` // 刷新时间
	List       []HotSearchSentence `json:"list"`        // 实时热点词
	DYError
}

// HotSearchSentencesRes 实时热点词
type HotSearchSentencesRes struct {
	Data  HotSearchSentencesData `json:"data"`
	Extra DYExtra                `json:"extra"`
}

// HotSearchSentences 获取实时热点词
func (m *Manager) HotSearchSentences(req HotSearchSentencesReq) (res HotSearchSentencesRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_HOT_SEARCH_SENTENCES, req.AccessToken), nil, nil)
	return res, err
}

// HotSearchTrendingSentencesReq 上升词请求
type HotSearchTrendingSentencesReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
}

// HotSearchTrendingSentence 上升词
type HotSearchTrendingSentence struct {
	HotLevel int64  `json:"hot_level"` // 热度 综合点赞、评论、转发等计算得出
	Sentence string `json:"sentence"`  // 热点词
	Label    int64  `json:"label"`     // 标签: * `0` - 无 * `1` - 新 * `2` - 推荐 * `3` - 热 * `4` - 爆 * `5` - 首发
}

// HotSearchTrendingSentencesData 上升词
type HotSearchTrendingSentencesData struct {
	List    []HotSearchTrendingSentence `json:"list"`     // 实时热点词
	Total   int32                       `json:"total"`    // 总数
	Cursor  int64                       `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool                        `json:"has_more"` // 更多数据
	DYError
}

// HotSearchTrendingSentencesRes 上升词
type HotSearchTrendingSentencesRes struct {
	Data  HotSearchTrendingSentencesData `json:"data"`
	Extra DYExtra                        `json:"extra"`
}

// HotSearchTrendingSentences 获取上升词
func (m *Manager) HotSearchTrendingSentences(req HotSearchTrendingSentencesReq) (res HotSearchTrendingSentencesRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&cursor=%d&count=%d", conf.API_HOT_SEARCH_TRENDING_SENTENCES, req.AccessToken, req.Cursor, req.Count), nil, nil)
	return res, err
}

// HotSearchVideosReq 热点词聚合的视频请求
type HotSearchVideosReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	HotSentence string // 热点词
}

// HotSearchVideosData 热点词聚合的视频
type HotSearchVideosData struct {
	List []Video `json:"list"`
	DYError
}

// HotSearchVideosRes 热点词聚合的视频
type HotSearchVideosRes struct {
	Data  HotSearchVideosData `json:"data"`
	Extra DYExtra             `json:"extra"`
}

// HotSearchVideos 获取热点词聚合的视频
func (m *Manager) HotSearchVideos(req HotSearchVideosReq) (res HotSearchVideosRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&hot_sentence=%s", conf.API_HOT_SEARCH_VIDEOS, req.AccessToken, req.HotSentence), nil, nil)
	return res, err
}
