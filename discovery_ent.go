package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// DiscoveryEntRankItemReq 抖音电影榜、抖音电视剧榜、抖音综艺榜请求
type DiscoveryEntRankItemReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Type        int32  // 榜单类型： * 1 - 电影 * 2 - 电视剧 * 3 - 综艺
	Version     int32  // 榜单版本：空值默认为本周榜单
}

// DiscoveryEntRankItemDataAlbum 抖音电影榜、抖音电视剧榜、抖音综艺榜
type DiscoveryEntRankItemDataAlbum struct {
	Id            string   `json:"id"`
	MaoyanId      string   `json:"maoyan_id,omitempty"`
	Name          string   `json:"name"`
	NameEn        string   `json:"name_en,omitempty"`
	ReleaseDate   string   `json:"release_date"`
	Directors     []string `json:"directors"`
	DiscussionHot int64    `json:"discussion_hot"`
	Poster        string   `json:"poster"`
	Tags          []string `json:"tags,omitempty"`
	Type          int32    `json:"type"`
	Actors        []string `json:"actors,omitempty"`
	Areas         []string `json:"areas,omitempty"`
	Hot           int64    `json:"hot"`
	SearchHot     int64    `json:"search_hot"`
	InfluenceHot  int64    `json:"influence_hot"`
	TopicHot      int64    `json:"topic_hot"`
}

// DiscoveryEntRankItemData 抖音电影榜、抖音电视剧榜、抖音综艺榜
type DiscoveryEntRankItemData struct {
	List []DiscoveryEntRankItemDataAlbum `json:"list"` // 实时热点词
	DYError
}

// DiscoveryEntRankItemRes 抖音电影榜、抖音电视剧榜、抖音综艺榜
type DiscoveryEntRankItemRes struct {
	Data  DiscoveryEntRankItemData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// DiscoveryEntRankItem 获取抖音电影榜、抖音电视剧榜、抖音综艺榜
func (m *Manager) DiscoveryEntRankItem(req DiscoveryEntRankItemReq) (res DiscoveryEntRankItemRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&type=%d&version=%d", conf.API_DISCOVERY_ENT_RANK_ITEM, req.AccessToken, req.Type, req.Version), nil, nil)
	return res, err
}

// DiscoveryEntRankVersionReq 抖音影视综榜单版本
type DiscoveryEntRankVersionReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	Type        int32  // 榜单类型： * 1 - 电影 * 2 - 电视剧 * 3 - 综艺
}

// DiscoveryEntRankVersionDataVersion 抖音影视综榜单版本
type DiscoveryEntRankVersionDataVersion struct {
	ActiveTime string `json:"active_time"` // 榜单生成时间
	EndTime    string `json:"end_time"`    // 榜单结束时间
	StartTime  string `json:"start_time"`  // 榜单起始时间
	Type       int32  `json:"type"`        // 榜单类型： * 1 - 电影 * 2 - 电视剧 * 3 - 综艺
	Version    int32  `json:"version"`     // 榜单版本
}

// DiscoveryEntRankVersionData 抖音影视综榜单版本
type DiscoveryEntRankVersionData struct {
	List    []DiscoveryEntRankVersionDataVersion `json:"list"`   // 榜单版本列表
	Cursor  int64                                `json:"cursor"` // 用于下一页请求的cursor
	HasMore bool                                 `json:"has_more"`
	DYError
}

// DiscoveryEntRankVersionRes 抖音影视综榜单版本
type DiscoveryEntRankVersionRes struct {
	Data  DiscoveryEntRankVersionData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// DiscoveryEntRankVersion 获取抖音影视综榜单版本
func (m *Manager) DiscoveryEntRankVersion(req DiscoveryEntRankVersionReq) (res DiscoveryEntRankVersionRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&cursor=%d&count=%d&type=%d", conf.API_DISCOVERY_ENT_RANK_VERSION, req.AccessToken, req.Cursor, req.Count, req.Type), nil, nil)
	return res, err
}
