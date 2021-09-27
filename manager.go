package douyingo

import (
	"fmt"
	"net/http"

	"github.com/zhangshuai/douyin-go/auth"
	"github.com/zhangshuai/douyin-go/client"
	"github.com/zhangshuai/douyin-go/conf"
)

// Manager Manager结构体
type Manager struct {
	client      *client.Client
	Credentials *auth.Credentials
}

// NewCredentials 获取认证
func NewCredentials(clientKey, clientSecret string) *auth.Credentials {
	return auth.New(clientKey, clientSecret)
}

// NewManager 创建新的Manager
func NewManager(credentials *auth.Credentials, tr http.RoundTripper) *Manager {
	client := client.DefaultClient
	client.Transport = newTransport(credentials, nil)
	return &Manager{
		client:      &client,
		Credentials: credentials,
	}
}

func (manager *Manager) url(format string, args ...interface{}) string {
	return conf.API_HTTP_SCHEME + conf.API_HOST + fmt.Sprintf(format, args...)
}
