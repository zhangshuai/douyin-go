package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// FansListReq 粉丝列表请求
type FansListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
}

// Fans 粉丝列表
type Fans struct {
	Nickname string `json:"nickname"` // 昵称
	Province string `json:"province"` // 省
	Avatar   string `json:"avatar"`   // 头像
	City     string `json:"city"`     // 城市
	Country  string `json:"country"`  // 国家
	Gender   int64  `json:"gender"`   // 性别: * `0` - 未知 * `1` - 男性 * `2` - 女性
	OpenId   string `json:"open_id"`  // 用户在当前应用的唯一标识
	UnionId  string `json:"union_id"` // 用户在当前开发者账号下的唯一标识（未绑定开发者账号没有该字段）
}

// FansListData 粉丝列表
type FansListData struct {
	List    []Fans `json:"list"`     // 粉丝列表
	Total   int64  `json:"total"`    // 粉丝总数
	Cursor  int64  `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool   `json:"has_more"` // 更多数据
	DYError
}

// FansListRes 粉丝列表
type FansListRes struct {
	Data  FansListData `json:"data"`
	Extra DYExtra      `json:"extra"`
}

// FansList 获取粉丝列表
func (m *Manager) FansList(req FansListReq) (res FansListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d", conf.API_FANS_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count), nil, nil)
	return res, err
}

// FansCheckReq 粉丝判断请求
type FansCheckReq struct {
	OpenId         string // 通过/oauth/access_token/获取，用户唯一标志
	FollowerOpenId string
	AccessToken    string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// FansCheckData 粉丝判断
type FansCheckData struct {
	IsFollower bool  `json:"is_follower"` // follower_open_id是否关注了open_id
	FollowTime int64 `json:"follow_time"` // 若关注了，则返回关注时间。单位为秒级时间戳
	DYError
}

// FansCheckRes 粉丝判断
type FansCheckRes struct {
	Data  FansCheckData `json:"data"`
	Extra DYExtra       `json:"extra"`
}

// FansCheck 获取粉丝判断
func (m *Manager) FansCheck(req FansCheckReq) (res FansCheckRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&follower_open_id=%s", conf.API_FANS_CHECK, req.AccessToken, req.OpenId, req.FollowerOpenId), nil, nil)
	return res, err
}
