package douyingo

import (
	"context"
	"net/url"

	"github.com/zhangshuai/douyin-go/conf"
)

// ItemCommentListReq 评论列表请求
type ItemCommentListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	ItemId      string // 视频id
}

// ItemComment 评论列表
type ItemComment struct {
	ReplyCommentTotal int32  `json:"reply_comment_total"` // 回复评论数
	Top               bool   `json:"top"`                 // 是否置顶评论
	CommentId         string `json:"comment_id"`          // 评论id
	CommentUserId     string `json:"comment_user_id"`     // 评论用户id
	Content           string `json:"content"`             // 评论内容
	CreateTime        int64  `json:"create_time"`         // 时间戳
	DiggCount         int32  `json:"digg_count"`          // 点赞数
}

// ItemCommentListData 评论列表
type ItemCommentListData struct {
	List    []ItemComment `json:"list"`     // 评论列表
	Cursor  int64         `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool          `json:"has_more"` // 更多数据
	DYError
}

// ItemCommentListRes 评论列表
type ItemCommentListRes struct {
	Data  ItemCommentListData `json:"data"`
	Extra DYExtra             `json:"extra"`
}

// ItemCommentList 获取评论列表
func (m *Manager) ItemCommentList(req ItemCommentListReq) (res ItemCommentListRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d&item_id=%s", conf.API_ITEM_COMMENT_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count, itemId), nil, nil)
	return res, err
}

// ItemCommentReplyListReq 评论回复列表请求
type ItemCommentReplyListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	ItemId      string // 视频id
	CommentId   string // 评论id
}

// ItemCommentReplyListData 评论回复列表
type ItemCommentReplyListData struct {
	List    []ItemComment `json:"list"`     // 评论回复列表
	Cursor  int64         `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool          `json:"has_more"` // 更多数据
	DYError
}

// ItemCommentReplyListRes 评论回复列表
type ItemCommentReplyListRes struct {
	Data  ItemCommentReplyListData `json:"data"`
	Extra DYExtra                  `json:"extra"`
}

// ItemCommentReplyList 获取评论回复列表
func (m *Manager) ItemCommentReplyList(req ItemCommentReplyListReq) (res ItemCommentReplyListRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	commentId := url.QueryEscape(req.CommentId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d&item_id=%s&comment_id=%s", conf.API_ITEM_COMMENT_REPLY_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count, itemId, commentId), nil, nil)
	return res, err
}

// ItemCommentReplyReq 回复视频评论请求
type ItemCommentReplyReq struct {
	OpenId      string               // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string               // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        ItemCommentReplyBody // 回复视频评论body
}

// ItemCommentReplyBody 回复视频评论
type ItemCommentReplyBody struct {
	CommentId string `json:"comment_id,omitempty"` // 需要回复的评论id（如果需要回复的是视频不传此字段）
	Content   string `json:"content"`              // 评论内容
	ItemId    string `json:"item_id"`              // 视频id
}

// ItemCommentReplyData 回复视频评论
type ItemCommentReplyData struct {
	CommentId string `json:"comment_id"` // 评论id
	DYError
}

// ItemCommentReplyRes 回复视频评论
type ItemCommentReplyRes struct {
	Data    ItemCommentReplyData `json:"data"`
	Extra   DYExtra              `json:"extra"`
	Message string               `json:"message"`
}

// ItemCommentReply 回复视频评论
func (m *Manager) ItemCommentReply(req ItemCommentReplyReq) (res ItemCommentReplyRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_ITEM_COMMENT_REPLY, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}
