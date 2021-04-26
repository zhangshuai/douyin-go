package douyinGo

import (
	"context"
)

type DataExternalBillboardReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Uri         string
}

type DataExternalBillboardItemVideo struct {
	ShareUrl  string `json:"share_url"`  // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	Title     string `json:"title"`      // 视频标题
	ItemCover string `json:"item_cover"` // 视频封面图
}

type DataExternalBillboardItem struct {
	Rank             int32                            `json:"rank"`                        // 排名
	RankChange       string                           `json:"rank_change,omitempty"`       // 排名变化, 如果上一期未上榜用-表示
	Nickname         string                           `json:"nickname"`                    // 昵称
	Avatar           string                           `json:"avatar"`                      // 头像
	FollowerCount    int64                            `json:"follower_count,omitempty"`    // 粉丝数
	OnbillbaordTimes int32                            `json:"onbillbaord_times,omitempty"` // 近一月在榜次数
	EffectValue      float64                          `json:"effect_value"`                // 影响力指数
	VideoList        []DataExternalBillboardItemVideo `json:"video_list,omitempty"`        // 视频列表
}

type DataExternalBillboardData struct {
	List []DataExternalBillboardItem `json:"list"`
	DYError
}

type DataExternalBillboardRes struct {
	Data  DataExternalBillboardData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// 获取榜单数据
func (m *Manager) DataExternalBillboard(req DataExternalBillboardReq) (res DataExternalBillboardRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", req.Uri, req.AccessToken), nil, nil)
	return res, err
}
