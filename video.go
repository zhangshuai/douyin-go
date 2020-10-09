package douyinGo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/zhangshuai/douyin-go/conf"
)

const (
	defaultWorkers   = 4                // 默认的并发上传的块数量
	defaultChunkSize = 20 * 1024 * 1024 // 默认的分片大小，20MB
)

type VideoListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
}

type Statistics struct {
	ShareCount    int `json:"share_count"`    // 分享数
	CommentCount  int `json:"comment_count"`  // 评论数
	DiggCount     int `json:"digg_count"`     // 点赞数
	DownloadCount int `json:"download_count"` // 下载数
	ForwardCount  int `json:"forward_count"`  // 转发数
	PlayCount     int `json:"play_count"`     // 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0。
}

type Video struct {
	Cover       string     `json:"cover"`        // 视频封面
	Statistics  Statistics `json:"statistics"`   // 统计数据
	Title       string     `json:"title"`        // 视频标题
	CreateTime  int64      `json:"create_time"`  // 视频创建时间戳
	IsReviewed  bool       `json:"is_reviewed"`  // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
	IsTop       bool       `json:"is_top"`       // 是否置顶
	ItemId      string     `json:"item_id"`      // 视频id
	ShareUrl    string     `json:"share_url"`    // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	VideoStatus int        `json:"video_status"` // 表示视频状态。1:已发布;2:不适宜公开;4:审核中
}

type VideoListResData struct {
	List    []Video `json:"list"`     // 视频列表
	Cursor  int64   `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool    `json:"has_more"` // 更多数据
	DYError
}

type VideoListRes struct {
	Data  VideoListResData `json:"data"`
	Extra DYExtra          `json:"extra"`
}

// 查询授权账号视频数据
func (m *Manager) VideoList(req VideoListReq) (res VideoListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d", conf.API_VIDEO_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count), nil, nil)
	return res, err
}

type VideoUploadReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	FilePath    string // 文件路径
}

type VideoUploadResVideo struct {
	Height  int64  `json:"height"`   // 视频高度
	Width   int64  `json:"width"`    // 视频宽度
	VideoId string `json:"video_id"` // 视频id
}

type VideoUploadResData struct {
	Video VideoUploadResVideo `json:"video,omitempty"`
	DYError
}

type VideoUploadRes struct {
	Data  VideoUploadResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// 上传视频到文件服务器
func (m *Manager) VideoUpload(req VideoUploadReq) (res *VideoUploadRes, err error) {
	f, err := os.Open(req.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return
	}
	fsize := fi.Size()
	fname := fi.Name()

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	head := make(textproto.MIMEHeader)
	extension := filepath.Ext(req.FilePath)
	head.Set("Content-Type", "video/"+strings.Replace(extension, ".", "", -1))
	head.Set("Content-Disposition", fmt.Sprintf(`form-data; name="video"; filename="%s"`, fname))
	if _, err := writer.CreatePart(head); err != nil {
		return nil, err
	}

	lastLine := fmt.Sprintf("\r\n--%s--\r\n", writer.Boundary())
	r := strings.NewReader(lastLine)

	bodyLen := int64(b.Len()) + fsize + int64(len(lastLine))
	mr := io.MultiReader(&b, f, r)
	contentType := writer.FormDataContentType()
	headers := http.Header{}
	headers.Add("Content-Type", contentType)
	err = m.client.CallWith64(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_UPLOAD, req.AccessToken, req.OpenId), headers, mr, bodyLen)
	return res, err
}

type VideoCreateReq struct {
	OpenId      string          // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoCreateBody // 创建视频body
}

type VideoCreateBody struct {
	TimelinessLabel   int64    `json:"timeliness_label,omitempty"`   // 时效新闻标签，1表示使用。暂不开放
	ArticlId          string   `json:"article_id,omitempty"`         // 文章ID，暂不开放
	MicroAppId        string   `json:"micro_app_id,omitempty"`       // 小程序id
	PoiId             string   `json:"poi_id,omitempty"`             // 地理位置id
	MicroAppUrl       string   `json:"micro_app_url,omitempty"`      // 吊起小程序时的参数
	ArticleTitle      string   `json:"article_title,omitempty"`      // 文章自定义标题，暂不开放
	CoverTsp          float64  `json:"cover_tsp,omitempty"`          // 将传入的指定时间点对应帧设置为视频封面（单位：秒）
	GameId            string   `json:"game_id,omitempty"`            // 游戏id。暂不开放
	MicroAppTitle     string   `json:"micro_app_title,omitempty"`    // 小程序标题
	PoiName           string   `json:"poi_name,omitempty"`           // 地理位置名称
	TimelinessKeyword string   `json:"timeliness_keyword,omitempty"` // 最多可添加3个，用`\|`隔开。暂不开放
	VideoId           string   `json:"video_id"`                     // video_id, 通过/video/upload/接口得到。注意每次调用/video/create/都要调用/video/upload/生成新的video_id。
	AtUsers           []string `json:"at_users,omitempty"`           // 如果需要at其他用户。将text中@nickname对应的open_id放到这里。
	GameContent       string   `json:"game_content,omitempty"`       // 游戏个性化参数
	Text              string   `json:"text,omitempty"`               // 视频标题， 可以带话题,@用户。 如title1#话题1 #话题2 @openid1 注意： 1. 话题审核依旧遵循抖音的审核逻辑，强烈建议第三方谨慎拟定话题名称，避免强导流行为。
}

type VideoCreateResData struct {
	ItemId string `json:"item_id"` // 视频id
	DYError
}

type VideoCreateRes struct {
	Data  VideoCreateResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// 创建抖音视频
func (m *Manager) VideoCreate(req VideoCreateReq) (res VideoCreateRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_CREATE, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

type VideoDeleteReq struct {
	OpenId      string          // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoDeleteBody // 删除视频body
}

type VideoDeleteBody struct {
	ItemId string `json:"item_id,omitempty"` // 抖音视频id
}

type VideoDeleteResData struct {
	DYError
}

type VideoDeleteRes struct {
	Data  VideoDeleteResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// 删除授权用户发布的视频
func (m *Manager) VideoDelete(req VideoDeleteReq) (res VideoDeleteRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_DELETE, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

type VideoDataReq struct {
	OpenId      string        // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string        // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoDataBody // 视频数据body
}

type VideoDataBody struct {
	ItemIds []string `json:"item_ids"` // item_id数组，仅能查询access_token对应用户上传的视频
}

type VideoDataResData struct {
	List []Video `json:"list"` // 视频数据列表
	DYError
}

type VideoDataRes struct {
	Data  VideoDataResData `json:"data"`
	Extra DYExtra          `json:"extra"`
}

// 查询指定视频数据
func (m *Manager) VideoData(req VideoDataReq) (res VideoDataRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_DATA, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

type VideoPartUploadInitReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type VideoPartUploadInitResData struct {
	UploadId string `json:"upload_id"` // 上传id
	DYError
}

type VideoPartUploadInitRes struct {
	Data  VideoPartUploadInitResData `json:"data"`
	Extra DYExtra                    `json:"extra"`
}

// 初始化上传
func (m *Manager) VideoPartUploadInit(req VideoPartUploadInitReq) (res VideoPartUploadInitRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_UPLOAD_PART_INIT, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}

type VideoPartUploadReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	FilePath    string // 文件路径
	UploadId    string // 分片上传的标记。有限时间为2小时。
	ChunkSize   int64  // 视频分片，建议20MB。但不能小于5MB。
	Workers     uint   // 并发上传的块数量
}

type VideoPartUploadData struct {
	DYError
}

type VideoPartUploadRes struct {
	Data  VideoPartUploadData `json:"data"`
	Extra DYExtra             `json:"extra"`
}

// 上传视频分片到文件服务器
func (m *Manager) VideoPartUpload(req VideoPartUploadReq) (res *VideoPartUploadRes, err error) {
	f, err := os.Open(req.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return
	}
	fsize := fi.Size()
	fname := fi.Name()

	var (
		workers   uint  = defaultWorkers
		chunkSize int64 = defaultChunkSize
		wg        sync.WaitGroup
	)
	tasks := make(chan bool, workers)

	if req.ChunkSize > 0 {
		chunkSize = req.ChunkSize
	}
	if req.Workers > 0 {
		workers = req.Workers
	}
	chunkTotal := int(fsize/int64(chunkSize) + 1)

	for i := 1; i <= chunkTotal; i++ {
		wg.Add(1)
		tasks <- true
		go func(i int) {
			defer func() {
				wg.Done()
				<-tasks
			}()
			var bodyLength int64 = chunkSize
			if i == chunkTotal {
				bodyLength = fsize % int64(chunkSize)
			}

			body := io.NewSectionReader(f, chunkSize*int64(i-1), int64(bodyLength))
			var b bytes.Buffer
			writer := multipart.NewWriter(&b)
			head := make(textproto.MIMEHeader)
			extension := filepath.Ext(req.FilePath)
			head.Set("Content-Type", "video/"+strings.Replace(extension, ".", "", -1))
			head.Set("Content-Disposition", fmt.Sprintf(`form-data; name="video"; filename="%s"`, fname))
			if _, err := writer.CreatePart(head); err != nil {
				return
			}

			lastLine := fmt.Sprintf("\r\n--%s--\r\n", writer.Boundary())
			r := strings.NewReader(lastLine)

			bodyLen := int64(b.Len()) + bodyLength + int64(len(lastLine))
			mr := io.MultiReader(&b, body, r)
			contentType := writer.FormDataContentType()
			headers := http.Header{}
			headers.Add("Content-Type", contentType)
			uploadId := url.QueryEscape(req.UploadId)
			err = m.client.CallWith64(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s&upload_id=%s&part_number=%d", conf.API_VIDEO_UPLOAD_PART_UPLOAD, req.AccessToken, req.OpenId, uploadId, i), headers, mr, bodyLen)
		}(i)
	}
	wg.Wait()
	return res, err
}

type VideoUploadPartCompleteReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	UploadId    string // 分片上传的标记。有限时间为2小时。
}

type VideoUploadPartCompleteRes struct {
	Data  VideoUploadResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// 完成上传
func (m *Manager) VideoUploadPartComplete(req VideoUploadPartCompleteReq) (res VideoUploadPartCompleteRes, err error) {
	uploadId := url.QueryEscape(req.UploadId)
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s&upload_id=%s", conf.API_VIDEO_UPLOAD_PART_COMPLETE, req.AccessToken, req.OpenId, uploadId), nil, nil)
	return res, err
}
