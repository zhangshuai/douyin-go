package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// DataExternalBillboardReq 榜单数据请求
type DataExternalBillboardReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Uri         string
}

// DataExternalBillboardItemVideo 榜单数据
type DataExternalBillboardItemVideo struct {
	ShareUrl  string `json:"share_url"`  // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	Title     string `json:"title"`      // 视频标题
	ItemCover string `json:"item_cover"` // 视频封面图
}

// DataExternalBillboardItem 榜单数据
type DataExternalBillboardItem struct {
	Rank             int32                            `json:"rank"`                        // 排名
	RankChange       string                           `json:"rank_change,omitempty"`       // 排名变化, 如果上一期未上榜用-表示
	Nickname         string                           `json:"nickname,omitempty"`          // 昵称
	Title            string                           `json:"title,omitempty"`             // 话题标题
	Avatar           string                           `json:"avatar,omitempty"`            // 头像
	FollowerCount    int64                            `json:"follower_count,omitempty"`    // 粉丝数
	OnbillbaordTimes int32                            `json:"onbillbaord_times,omitempty"` // 近一月在榜次数
	EffectValue      float64                          `json:"effect_value"`                // 影响力指数
	VideoList        []DataExternalBillboardItemVideo `json:"video_list,omitempty"`        // 视频列表
}

// DataExternalBillboardData 榜单数据
type DataExternalBillboardData struct {
	List []DataExternalBillboardItem `json:"list"`
	DYError
}

// DataExternalBillboardRes 榜单数据
type DataExternalBillboardRes struct {
	Data  DataExternalBillboardData `json:"data"`
	Extra DYExtra                   `json:"extra"`
}

// DataExternalBillboard 获取榜单数据
func (m *Manager) DataExternalBillboard(req DataExternalBillboardReq) (res DataExternalBillboardRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", req.Uri, req.AccessToken), nil, nil)
	return res, err
}

// DataExternalBillboardPropReq 道具榜单数据请求
type DataExternalBillboardPropReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// DataExternalBillboardPropItem 道具榜单数据
type DataExternalBillboardPropItem struct {
	Rank               int32   `json:"rank"`                 // 排名
	RankChange         string  `json:"rank_change"`          // 排名变化
	Name               string  `json:"name"`                 // 道具名
	ShowCnt            float64 `json:"show_cnt"`             // 展现量，离线数据（统计前一日数据）
	ShootCnt           float64 `json:"shoot_cnt"`            // 开拍量，离线数据（统计前一日数据）
	DailyIssueCnt      float64 `json:"daily_issue_cnt"`      // 日投稿量，离线数据（统计前一日数据）
	DailyIssuePercent  string  `json:"daily_issue_percent"`  // 日投稿占比，格式:XX.XX% 精确小数点后2位 离线数据（统计前一日数据）
	DailyCollectionCnt float64 `json:"daily_collection_cnt"` // 日收藏数，离线数据（统计前一日数据）
	DailyPlayCnt       float64 `json:"daily_play_cnt"`       // 日播放数，离线数据（统计前一日数据）
	EffectValue        float64 `json:"effect_value"`         // 影响力指数
}

// DataExternalBillboardPropData 道具榜单数据
type DataExternalBillboardPropData struct {
	List []DataExternalBillboardPropItem `json:"list"`
	DYError
}

// DataExternalBillboardPropRes 道具榜单数据
type DataExternalBillboardPropRes struct {
	Data  DataExternalBillboardPropData `json:"data"`
	Extra DYExtra                       `json:"extra"`
}

// DataExternalBillboardProp 获取道具榜单数据
func (m *Manager) DataExternalBillboardProp(req DataExternalBillboardPropReq) (res DataExternalBillboardPropRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_DATA_EXTERNAL_BILLBOARD_PROP, req.AccessToken), nil, nil)
	return res, err
}

// DataExternalBillboardHotVideoReq 热门视频数据请求
type DataExternalBillboardHotVideoReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// DataExternalBillboardHotVideoItem 热门视频数据
type DataExternalBillboardHotVideoItem struct {
	Rank         int32   `json:"rank"`          // 排名
	ShareUrl     string  `json:"share_url"`     // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	Title        string  `json:"title"`         // 视频标题
	Author       string  `json:"author"`        // 视频发布者
	PlayCount    int64   `json:"play_count"`    // 播放数，离线数据（统计前一日数据）
	DiggCount    int64   `json:"digg_count"`    // 点赞数，离线数据（统计前一日数据）
	CommentCount int64   `json:"comment_count"` // 评论数，离线数据（统计前一日数据）
	HotWords     string  `json:"hot_words"`     // 视频热词（以,隔开）
	HotValue     float64 `json:"hot_value"`     // 热度指数
	ItemCover    string  `json:"item_cover"`    // 视频封面图
}

// DataExternalBillboardHotVideoData 热门视频数据
type DataExternalBillboardHotVideoData struct {
	List []DataExternalBillboardHotVideoItem `json:"list"`
	DYError
}

// DataExternalBillboardHotVideoRes 热门视频数据
type DataExternalBillboardHotVideoRes struct {
	Data  DataExternalBillboardHotVideoData `json:"data"`
	Extra DYExtra                           `json:"extra"`
}

// DataExternalBillboardHotVideo 获取热门视频数据
func (m *Manager) DataExternalBillboardHotVideo(req DataExternalBillboardHotVideoReq) (res DataExternalBillboardHotVideoRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_DATA_EXTERNAL_BILLBOARD_HOT_VIDEO, req.AccessToken), nil, nil)
	return res, err
}

// DataExternalBillboardLiveReq 直播榜数据请求
type DataExternalBillboardLiveReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
}

// DataExternalBillboardLiveItem 直播榜数据
type DataExternalBillboardLiveItem struct {
	Rank     int32   `json:"rank"`      // 排名
	Cover    string  `json:"cover"`     // 直播封面
	Title    string  `json:"title"`     // 直播标题
	Nickname string  `json:"nickname"`  // 昵称
	HotValue float64 `json:"hot_value"` // 热度指数
}

// DataExternalBillboardLiveData 直播榜数据
type DataExternalBillboardLiveData struct {
	List []DataExternalBillboardLiveItem `json:"list"`
	DYError
}

// DataExternalBillboardLiveRes 直播榜数据
type DataExternalBillboardLiveRes struct {
	Data  DataExternalBillboardLiveData `json:"data"`
	Extra DYExtra                       `json:"extra"`
}

// DataExternalBillboardLive 获取直播榜数据
func (m *Manager) DataExternalBillboardLive(req DataExternalBillboardLiveReq) (res DataExternalBillboardLiveRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", conf.API_DATA_EXTERNAL_BILLBOARD_LIVE, req.AccessToken), nil, nil)
	return res, err
}

// DataExternalBillboardMusicReq 音乐榜单数据请求
type DataExternalBillboardMusicReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Uri         string
}

// DataExternalBillboardMusicItem 音乐榜单数据
type DataExternalBillboardMusicItem struct {
	Rank     int32  `json:"rank"`      // 排名
	Cover    string `json:"cover"`     // 音乐封面
	Title    string `json:"title"`     // 歌曲标题
	Duration int32  `json:"duration"`  // 时长，精确到秒
	Author   string `json:"author"`    // 作者昵称
	UseCount int64  `json:"use_count"` // 使用量
	ShareUrl string `json:"share_url"` // 音乐分享链接
}

// DataExternalBillboardMusicData 音乐榜单数据
type DataExternalBillboardMusicData struct {
	List []DataExternalBillboardMusicItem `json:"list"`
	DYError
}

// DataExternalBillboardMusicRes 音乐榜单数据
type DataExternalBillboardMusicRes struct {
	Data  DataExternalBillboardMusicData `json:"data"`
	Extra DYExtra                        `json:"extra"`
}

// DataExternalBillboardMusic 获取音乐榜单数据
func (m *Manager) DataExternalBillboardMusic(req DataExternalBillboardMusicReq) (res DataExternalBillboardMusicRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s", req.Uri, req.AccessToken), nil, nil)
	return res, err
}
