package douyingo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangshuai/douyin-go/conf"
)

// ImageUploadReq 上传图片到文件服务器请求
type ImageUploadReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	FilePath    string // 文件路径
}

// ImageUploadResImage 上传图片到文件服务器
type ImageUploadResImage struct {
	Height  int64  `json:"height"`   // 图片高度
	Width   int64  `json:"width"`    // 图片宽度
	ImageId string `json:"image_id"` // 图片id
}

// ImageUploadResData 上传图片到文件服务器
type ImageUploadResData struct {
	Image ImageUploadResImage `json:"image,omitempty"`
	DYError
}

// ImageUploadRes 上传图片到文件服务器
type ImageUploadRes struct {
	Data  ImageUploadResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// ImageUpload 上传图片到文件服务器
func (m *Manager) ImageUpload(req ImageUploadReq) (res *ImageUploadRes, err error) {
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
	head.Set("Content-Type", "image/"+strings.Replace(extension, ".", "", -1))
	head.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, fname))
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
	err = m.client.CallWith64(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_IMAGE_UPLOAD, req.AccessToken, req.OpenId), headers, mr, bodyLen)
	return res, err
}

// ImageCreateReq 发布图片请求
type ImageCreateReq struct {
	OpenId      string          // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        ImageCreateBody // 请求body
}

// ImageCreateBody 发布图片
type ImageCreateBody struct {
	PoiId         string   `json:"poi_id,omitempty"`          // 地理位置id
	PoiName       string   `json:"poi_name,omitempty"`        // 地理位置名称
	Text          string   `json:"text,omitempty"`            // 标题，可以带话题。 如title1#话题1 #话题2 注意：话题审核依旧遵循抖音的审核逻辑，强烈建议第三方谨慎拟定话题名称，避免强导流行为。
	MicroAppId    string   `json:"micro_app_id,omitempty"`    // 小程序id
	MicroAppTitle string   `json:"micro_app_title,omitempty"` // 小程序标题
	MicroAppUrl   string   `json:"micro_app_url,omitempty"`   // 吊起小程序时的参数
	ImageId       string   `json:"image_id"`                  // 通过/image/upload/接口得到。
	AtUsers       []string `json:"at_users,omitempty"`        // 如果需要at其他用户。将text中@nickname对应的open_id放到这里。
}

// ImageCreateResData 发布图片
type ImageCreateResData struct {
	ItemId string `json:"item_id"` // 抖音图片id
	DYError
}

// ImageCreateRes 发布图片
type ImageCreateRes struct {
	Data  ImageCreateResData `json:"data"`
	Extra DYExtra            `json:"extra"`
}

// ImageCreate 发布图片
func (m *Manager) ImageCreate(req ImageCreateReq) (res *ImageCreateRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_IMAGE_CREATE, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}
