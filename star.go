package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// StarHotListReq 抖音星图达人热榜请求
type StarHotListReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	HotListType int64  // 达人热榜类型 * `1` - 星图指数榜 * `2` - 涨粉指数榜 * `3` - 性价比指数榜 * `4` - 种草指数榜 * `5` - 精选指数榜 * `6` - 传播指数榜
}

// StarHotList 抖音星图达人热榜
type StarHotList struct {
	Follower int64    `json:"follower"`  // 粉丝数
	NickName string   `json:"nick_name"` // 达人昵称
	Rank     int64    `json:"rank"`      // 热榜排名
	Score    float64  `json:"score"`     // 热榜类型对应的热榜指数
	Tags     []string `json:"tags"`      // 标签
	UniqueId string   `json:"unique_id"` // 抖音号
}

// StarHotListData 抖音星图达人热榜
type StarHotListData struct {
	HotListType            int64         `json:"hot_list_type"`             // 刷新时间
	HotListUpdateTimestamp int64         `json:"hot_list_update_timestamp"` // 达人热榜更新时间戳
	HotListDescription     string        `json:"hot_list_description"`      // 热榜类型说明 【筛选规则】：综合评估创作者近期作品的有效视频数据、性价比、信用分、有效涨粉、种草指数等加权取值进行排序，以内容数据为主要衡量标准。【达人优势】：综合能力高，具有较大的商业价值。【适用场景】：适用于大部分营销场景，可作为各领域商业投放的必选名单。
	List                   []StarHotList `json:"list"`                      // 实时热点词
	DYError
}

// StarHotListRes 抖音星图达人热榜
type StarHotListRes struct {
	Data  StarHotListData `json:"data"`
	Extra DYExtra         `json:"extra"`
}

// StarHotList 获取抖音星图达人热榜
func (m *Manager) StarHotList(req StarHotListReq) (res StarHotListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&hot_list_type=%d", conf.API_STAR_HOT_LIST, req.AccessToken, req.HotListType), nil, nil)
	return res, err
}

// StarAuthorScoreReq 抖音星图达人指数请求
type StarAuthorScoreReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
}

// StarAuthorScoreData 抖音星图达人指数
type StarAuthorScoreData struct {
	ShopScore        float64 `json:"shop_score"`         // 种草指数
	StarScore        float64 `json:"star_score"`         // 星图指数
	UniqueId         string  `json:"unique_id"`          // 达人抖音号
	UpdateTimestamp  int64   `json:"update_timestamp"`   // 达人指数更新时间戳
	CooperationScore float64 `json:"cooperation_score"`  // 合作指数
	CpScore          float64 `json:"cp_score"`           // 性价比指数
	NickName         string  `json:"nick_name"`          // 达人昵称
	GrowthScore      float64 `json:"growth_score"`       // 涨粉指数
	SpreadScore      float64 `json:"spread_score"`       // 传播指数
	Follower         int64   `json:"follower,omitempty"` // 粉丝数
	DYError
}

// StarAuthorScoreRes 抖音星图达人指数
type StarAuthorScoreRes struct {
	Data  StarAuthorScoreData `json:"data"`
	Extra DYExtra             `json:"extra"`
}

// StarAuthorScore 获取抖音星图达人指数
func (m *Manager) StarAuthorScore(req StarAuthorScoreReq) (res StarAuthorScoreRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s", conf.API_STAR_AUTHOR_SCORE, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}

// StarAuthorScoreV2Req 抖音星图达人指数数据V2
type StarAuthorScoreV2Req struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	UniqueId    string // 达人抖音号
}

// StarAuthorScoreV2Res 抖音星图达人指数数据V2
type StarAuthorScoreV2Res struct {
	Data  StarAuthorScoreData `json:"data"`
	Extra DYExtra             `json:"extra"`
}

// StarAuthorScoreV2 获取抖音星图达人指数数据V2
func (m *Manager) StarAuthorScoreV2(req StarAuthorScoreV2Req) (res StarAuthorScoreV2Res, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&unique_id=%s", conf.API_STAR_AUTHOR_SCORE_V2, req.AccessToken, req.UniqueId), nil, nil)
	return res, err
}
