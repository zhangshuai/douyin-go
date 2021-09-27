package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// FollowingListReq 关注列表请求
type FollowingListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
}

// Following 关注列表
type Following struct {
	Nickname string `json:"nickname"` // 昵称
	Province string `json:"province"` // 省
	Avatar   string `json:"avatar"`   // 头像
	City     string `json:"city"`     // 城市
	Country  string `json:"country"`  // 国家
	Gender   int64  `json:"gender"`   // 性别: * `0` - 未知 * `1` - 男性 * `2` - 女性
	OpenId   string `json:"open_id"`  // 用户在当前应用的唯一标识
	UnionId  string `json:"union_id"` // 用户在当前开发者账号下的唯一标识（未绑定开发者账号没有该字段）
}

// FollowingListData 关注列表
type FollowingListData struct {
	List    []Following `json:"list"`     // 关注列表
	Cursor  int64       `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool        `json:"has_more"` // 更多数据
	DYError
}

// FollowingListRes 关注列表
type FollowingListRes struct {
	Data  FollowingListData `json:"data"`
	Extra DYExtra           `json:"extra"`
}

// FollowingList 获取关注列表
func (m *Manager) FollowingList(req FollowingListReq) (res FollowingListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d", conf.API_FOLLOWING_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count), nil, nil)
	return res, err
}
