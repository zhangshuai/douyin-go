package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// VideoSearchReq 关键词视频搜索请求
type VideoSearchReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	Keyword     string // 关键词
}

// VideoSearchResult 关键词视频搜索
type VideoSearchResult struct {
	Cover      string     `json:"cover"`       // 视频封面
	Statistics Statistics `json:"statistics"`  // 统计数据
	Title      string     `json:"title"`       // 视频标题
	CreateTime int64      `json:"create_time"` // 视频创建时间戳
	IsReviewed bool       `json:"is_reviewed"` // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
	IsTop      bool       `json:"is_top"`      // 是否置顶
	ItemId     string     `json:"item_id"`     // 视频id
	ShareUrl   string     `json:"share_url"`   // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	Nickname   string     `json:"nickname"`    // 昵称
	OpenId     string     `json:"open_id"`     // 作者openID
	Avatar     string     `json:"avatar"`      // 头像
	SecItemId  string     `json:"sec_item_id"` // 特殊加密的视频id通过用户视频搜索的评论接口获取到
}

// VideoSearchData 关键词视频搜索
type VideoSearchData struct {
	List    []VideoSearchResult `json:"list"`     // 由于置顶的原因, list长度可能比count指定的数量多一些或少一些。
	Cursor  int64               `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool                `json:"has_more"` // 更多数据
	DYError
}

// VideoSearchRes 关键词视频搜索
type VideoSearchRes struct {
	Data  VideoSearchData `json:"data"`
	Extra DYExtra         `json:"extra"`
}

// VideoSearch 关键词视频搜索
func (m *Manager) VideoSearch(req VideoSearchReq) (res VideoSearchRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d&keyword=%s", conf.API_VIDEO_SEARCH, req.AccessToken, req.OpenId, req.Cursor, req.Count, req.Keyword), nil, nil)
	return res, err
}

// VideoSearchCommentListReq 关键词视频评论列表请求
type VideoSearchCommentListReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	SecItemId   string // 视频搜索接口返回的加密的视频id
}

// VideoSearchCommentListData 关键词视频评论列表
type VideoSearchCommentListData struct {
	List    []ItemComment `json:"list"`
	Cursor  int64         `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool          `json:"has_more"` // 更多数据
	DYError
}

// VideoSearchCommentListRes 关键词视频评论列表
type VideoSearchCommentListRes struct {
	Data  VideoSearchCommentListData `json:"data"`
	Extra DYExtra                    `json:"extra"`
}

// VideoSearchCommentList 关键词视频评论列表
func (m *Manager) VideoSearchCommentList(req VideoSearchCommentListReq) (res VideoSearchCommentListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&cursor=%d&count=%d&sec_item_id=%s", conf.API_VIDEO_SEARCH_COMMENT_LIST, req.AccessToken, req.Cursor, req.Count, req.SecItemId), nil, nil)
	return res, err
}

// VideoSearchCommentReplyListReq 关键词视频评论回复列表请求
type VideoSearchCommentReplyListReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	SecItemId   string // 视频搜索接口返回的加密的视频id
	CommentId   string // 评论id
}

// VideoSearchCommentReplyListData 关键词视频评论回复列表
type VideoSearchCommentReplyListData struct {
	List    []ItemComment `json:"list"`
	Cursor  int64         `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool          `json:"has_more"` // 更多数据
	DYError
}

// VideoSearchCommentReplyListRes 关键词视频评论回复列表
type VideoSearchCommentReplyListRes struct {
	Data  VideoSearchCommentReplyListData `json:"data"`
	Extra DYExtra                         `json:"extra"`
}

// VideoSearchCommentReplyList 关键词视频评论回复列表
func (m *Manager) VideoSearchCommentReplyList(req VideoSearchCommentReplyListReq) (res VideoSearchCommentReplyListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&cursor=%d&count=%d&sec_item_id=%s&comment_id=%s", conf.API_VIDEO_SEARCH_COMMENT_REPLY_LIST, req.AccessToken, req.Cursor, req.Count, req.SecItemId, req.CommentId), nil, nil)
	return res, err
}

// VideoSearchCommentReplyReq 关键词视频评论回复请求
type VideoSearchCommentReplyReq struct {
	OpenId      string                      // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string                      // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoSearchCommentReplyBody // 回复视频评论body
}

// VideoSearchCommentReplyBody 关键词视频评论回复
type VideoSearchCommentReplyBody struct {
	CommentId string `json:"comment_id,omitempty"` // 需要回复的评论id（如果需要回复的是视频不传此字段）
	Content   string `json:"content"`              // 评论内容
	SecItemId string `json:"sec_item_id"`          // 视频搜索接口返回的加密的视频id
}

// VideoSearchCommentReplyData 关键词视频评论回复
type VideoSearchCommentReplyData struct {
	CommentId string `json:"comment_id"` // 评论id
	DYError
}

// VideoSearchCommentReplyRes 关键词视频评论回复
type VideoSearchCommentReplyRes struct {
	Data    VideoSearchCommentReplyData `json:"data"`
	Extra   DYExtra                     `json:"extra"`
	Message string                      `json:"message"`
}

// VideoSearchCommentReply 关键词视频评论回复
func (m *Manager) VideoSearchCommentReply(req VideoSearchCommentReplyReq) (res VideoSearchCommentReplyRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_SEARCH_COMMENT_REPLY, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}
