package douyinGo

import (
	"fmt"
	"net/http"

	"github.com/zhangshuai/douyin-go/auth"
	"github.com/zhangshuai/douyin-go/client"
	"github.com/zhangshuai/douyin-go/conf"
)

type Manager struct {
	client      *client.Client
	Credentials *auth.Credentials
}

func NewCredentials(clientKey, clientSecret string) *auth.Credentials {
	return auth.New(clientKey, clientSecret)
}

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
