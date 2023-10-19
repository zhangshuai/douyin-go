package douyingo

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

// VideoListReq 视频列表请求
type VideoListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
}

// Statistics 视频列表
type Statistics struct {
	ShareCount    int `json:"share_count"`    // 分享数
	CommentCount  int `json:"comment_count"`  // 评论数
	DiggCount     int `json:"digg_count"`     // 点赞数
	DownloadCount int `json:"download_count"` // 下载数
	ForwardCount  int `json:"forward_count"`  // 转发数
	PlayCount     int `json:"play_count"`     // 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0。
}

// Video 视频列表
type Video struct {
	Cover       string     `json:"cover"`        // 视频封面
	Statistics  Statistics `json:"statistics"`   // 统计数据
	Title       string     `json:"title"`        // 视频标题
	CreateTime  int64      `json:"create_time"`  // 视频创建时间戳
	IsReviewed  bool       `json:"is_reviewed"`  // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
	IsTop       bool       `json:"is_top"`       // 是否置顶
	ItemId      string     `json:"item_id"`      // 视频id
	ShareUrl    string     `json:"share_url"`    // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	VideoStatus int        `json:"video_status"` // 表示视频状态。1:已发布;2:不适宜公开;4:审核中;5:公开视频
	MediaType   int        `json:"media_type"`   // 媒体类型。2:图集;4:视频
	VideoId     string     `json:"video_id"`     // 视频真实 ID
}

// VideoListResData 视频列表
type VideoListResData struct {
	List    []Video `json:"list"`     // 视频列表
	Cursor  int64   `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool    `json:"has_more"` // 更多数据
	DYError
}

// VideoListRes 视频列表
type VideoListRes struct {
	Data  VideoListResData `json:"data"`
	Extra DYExtra          `json:"extra"`
}

// VideoList 查询授权账号视频数据
func (m *Manager) VideoList(req VideoListReq) (res VideoListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d", conf.API_VIDEO_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count), nil, nil)
	return res, err
}

// VideoUploadReq 上传视频到文件服务器
type VideoUploadReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	FilePath    string // 文件路径
}

// VideoUploadResVideo 上传视频到文件服务器
type VideoUploadResVideo struct {
	Height  int64  `json:"height"`   // 视频高度
	Width   int64  `json:"width"`    // 视频宽度
	VideoId string `json:"video_id"` // 视频id
}

// VideoUploadResData 上传视频到文件服务器
type VideoUploadResData struct {
	Video VideoUploadResVideo `json:"video,omitempty"`
	DYError
}

// VideoUploadRes 上传视频到文件服务器
type VideoUploadRes struct {
	Data  VideoUploadResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// VideoUpload 上传视频到文件服务器
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

// VideoCreateReq 创建抖音视频请求
type VideoCreateReq struct {
	OpenId      string          // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoCreateBody // 创建视频body
}

// VideoCreateBody 创建抖音视频
type VideoCreateBody struct {
	VideoId             string   `json:"video_id"`                         // 必须上传加密的 video_id。加密的 video_id 通过调用 /video/upload 接口可以得到。
	Text                string   `json:"text,omitempty"`                   // 视频标题。可以带话题，@用户。注意：话题审核依旧遵循抖音的审核逻辑，强烈建议第三方谨慎拟定话题名称，避免强导流行为。
	AtUsers             []string `json:"at_users,omitempty"`               // 如果需要 at 其他用户。将 text 中 @nickname 对应的 open_id 放到这里。
	MicroAppTitle       string   `json:"micro_app_title,omitempty"`        // 小程序标题
	MicroAppUrl         string   `json:"micro_app_url,omitempty"`          // 开发者在小程序中生成该页面时写的 path 地址
	MicroAppId          string   `json:"micro_app_id,omitempty"`           // 小程序 id
	PoiId               string   `json:"poi_id,omitempty"`                 // 地理位置 id，poi_id 可通过"查询视频携带的地点信息"能力获取
	CoverTsp            float64  `json:"cover_tsp,omitempty"`              // 将传入的指定时间点对应帧设置为视频封面（单位：秒）
	PoiCommerce         bool     `json:"poi_commerce,omitempty"`           // 为 true 时，如果用户拥有门店推广能力，则用户发布视频所添加的地理位置默认开启门店推广
	CustomCoverImageUrl string   `json:"custom_cover_image_url,omitempty"` // 自定义封面图片，参数为图片上传接口返回的 image_id
}

// VideoCreateResData 创建抖音视频
type VideoCreateResData struct {
	ItemId string `json:"item_id"` // 视频id
	DYError
}

// VideoCreateRes 创建抖音视频
type VideoCreateRes struct {
	Data  VideoCreateResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// VideoCreate 创建抖音视频
func (m *Manager) VideoCreate(req VideoCreateReq) (res VideoCreateRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_CREATE, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

// VideoDeleteReq 删除视频请求
type VideoDeleteReq struct {
	OpenId      string          // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoDeleteBody // 删除视频body
}

// VideoDeleteBody 删除视频
type VideoDeleteBody struct {
	ItemId string `json:"item_id,omitempty"` // 抖音视频id
}

// VideoDeleteResData 删除视频
type VideoDeleteResData struct {
	DYError
}

// VideoDeleteRes 删除视频
type VideoDeleteRes struct {
	Data  VideoDeleteResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// VideoDelete 删除授权用户发布的视频
func (m *Manager) VideoDelete(req VideoDeleteReq) (res VideoDeleteRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_DELETE, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

// VideoDataReq 视频数据请求
type VideoDataReq struct {
	OpenId      string        // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string        // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoDataBody // 视频数据body
}

// VideoDataBody 视频数据
type VideoDataBody struct {
	ItemIds  []string `json:"item_ids,omitempty"`  // item_id 数组，仅能查询 access_token 对应用户上传的视频（与video_ids字段二选一，平台优先处理item_ids）
	VideoIds []string `json:"video_ids,omitempty"` // video_id 数组，仅能查询 access_token 对应用户上传的视频（与item_ids字段二选一，平台优先处理item_ids）
}

// VideoDataResData 视频数据
type VideoDataResData struct {
	List []Video `json:"list"` // 视频数据列表
	DYError
}

// VideoDataRes 视频数据
type VideoDataRes struct {
	Data  VideoDataResData `json:"data"`
	Extra DYExtra          `json:"extra"`
}

// VideoData 查询指定视频数据
func (m *Manager) VideoData(req VideoDataReq) (res VideoDataRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_DATA, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

// VideoPartUploadInitReq 初始化分片上传请求
type VideoPartUploadInitReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// VideoPartUploadInitResData 初始化分片上传
type VideoPartUploadInitResData struct {
	UploadId string `json:"upload_id,omitempty"` // 上传id
	DYError
}

// VideoPartUploadInitRes 初始化分片上传
type VideoPartUploadInitRes struct {
	Data  VideoPartUploadInitResData `json:"data"`
	Extra DYExtra                    `json:"extra"`
}

// VideoPartUploadInit 初始化分片上传
func (m *Manager) VideoPartUploadInit(req VideoPartUploadInitReq) (res VideoPartUploadInitRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_UPLOAD_PART_INIT, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}

// VideoPartUploadReq 上传视频分片到文件服务器请求
type VideoPartUploadReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	FilePath    string // 文件路径
	UploadId    string // 分片上传的标记。有限时间为2小时。
	ChunkSize   int64  // 视频分片，建议20MB。但不能小于5MB。
	Workers     uint   // 并发上传的块数量
}

// VideoPartUploadData 上传视频分片到文件服务器
type VideoPartUploadData struct {
	DYError
}

// VideoPartUploadRes 上传视频分片到文件服务器
type VideoPartUploadRes struct {
	Data  VideoPartUploadData `json:"data"`
	Extra DYExtra             `json:"extra"`
}

// VideoPartUpload 上传视频分片到文件服务器
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

	if req.ChunkSize > 0 {
		chunkSize = req.ChunkSize
	}
	if req.Workers > 0 {
		workers = req.Workers
	}
	tasks := make(chan bool, workers)
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

// VideoUploadPartCompleteReq 分片上传视频请求
type VideoUploadPartCompleteReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	UploadId    string // 分片上传的标记。有限时间为2小时。
}

// VideoUploadPartCompleteRes 分片上传视频
type VideoUploadPartCompleteRes struct {
	Data  VideoUploadResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// VideoUploadPartComplete 完成分片上传视频
func (m *Manager) VideoUploadPartComplete(req VideoUploadPartCompleteReq) (res VideoUploadPartCompleteRes, err error) {
	uploadId := url.QueryEscape(req.UploadId)
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s&upload_id=%s", conf.API_VIDEO_UPLOAD_PART_COMPLETE, req.AccessToken, req.OpenId, uploadId), nil, nil)
	return res, err
}

// VideoGetIframeData 获取IFrame代码响应数据
type VideoGetIframeData struct {
	IframeCode  string `json:"iframe_code"`  // 返回IFrame代码片段，用于直接注入页面
	VideoTitle  string `json:"video_title"`  // 视频标题
	VideoWidth  int64  `json:"video_width"`  // 视频宽度
	VideoHeight int64  `json:"video_height"` // 视频高度
}

// VideoGetIframeByVideoReq 通过VideoID获取IFrame代码
type VideoGetIframeByVideoReq struct {
	VideoId string // 视频VideoID
}

// VideoGetIframeByVideoReq 通过VideoID获取IFrame代码响应
type VideoGetIframeByVideoRes struct {
	Data VideoGetIframeData `json:"data"`
	DYExtraV1
}

// VideoGetIframeByVideo 通过VideoID获取IFrame代码
func (m *Manager) VideoGetIframeByVideo(req VideoGetIframeByVideoReq) (res VideoGetIframeByVideoRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?video_id=%s", conf.API_VIDEO_GET_IFRAME_BY_VIDEO, req.VideoId), nil, nil)
	return res, err
}

// VideoGetIframeByItemReq 通过ItemID获取IFrame代码
type VideoGetIframeByItemReq struct {
	ClientKey string // 应用标识
	ItemId    string // 视频ItemID
}

// VideoGetIframeByItemReq 通过ItemID获取IFrame代码响应
type VideoGetIframeByItemRes struct {
	Data VideoGetIframeData `json:"data"`
	DYExtraV1
}

// VideoGetIframeByItem 通过ItemID获取IFrame代码
func (m *Manager) VideoGetIframeByItem(req VideoGetIframeByItemReq) (res VideoGetIframeByItemRes, err error) {
	// itemID需要Base64URL编码
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?client_key=%s&item_id=%s", conf.API_VIDEO_GET_IFRAME_BY_ITEM, req.ClientKey, req.ItemId), nil, nil)
	return res, err
}
