package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// OauthParam 授权参数
type OauthParam struct {
	Scope         string // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	OptionalScope string // 应用授权可选作用域,多个授权作用域以英文逗号（,）分隔，每一个授权作用域后需要加上一个是否默认勾选的参数，1为默认勾选，0为默认不勾选
	RedirectUri   string // 授权成功后的回调地址，必须以http/https开头。域名必须对应申请应用时填写的域名，如不清楚请联系应用申请人。
	State         string // 用于保持请求和回调的状态
}

// OauthConnect 生成授权链接,获取授权码
func (m *Manager) OauthConnect(param OauthParam) string {
	return m.url("%s?client_key=%s&response_type=code&scope=%s&optionalScope=%s&redirect_uri=%s&state=%s", conf.API_OAUTH_CONNECT, m.Credentials.ClientKey, param.Scope, param.OptionalScope, param.RedirectUri, param.State)
}

// OauthAccessTokenReq access_token请求
type OauthAccessTokenReq struct {
	Code string // 授权码
}

// OauthAccessTokenResData access_token
type OauthAccessTokenResData struct {
	AccessToken  string `json:"access_token"`  // 接口调用凭证
	UnionId      string `json:"union_id"`      // 当且仅当该网站应用已获得该用户的userinfo授权时，才会出现该字段。
	Scope        string `json:"scope"`         // 用户授权的作用域(Scope)，使用逗号（,）分隔，开放平台几乎几乎每个接口都需要特定的Scope。
	ExpiresIn    uint64 `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	OpenId       string `json:"open_id"`       // 授权用户唯一标识
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	DYError
}

// OauthAccessTokenRes access_token
type OauthAccessTokenRes struct {
	Data    OauthAccessTokenResData `json:"data"`
	Message string                  `json:"message"`
}

// OauthAccessToken 获取access_token
func (m *Manager) OauthAccessToken(req OauthAccessTokenReq) (res OauthAccessTokenRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?grant_type=authorization_code&client_key=%s&client_secret=%s&code=%s", conf.API_OAUTH_ACCESS_TOKEN, m.Credentials.ClientKey, m.Credentials.ClientSecret, req.Code), nil, nil)
	return res, err
}

// OauthClientAccessTokenResData client_token
type OauthClientAccessTokenResData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint64 `json:"expires_in"`
	DYError
}

// OauthClientAccessTokenRes client_token
type OauthClientAccessTokenRes struct {
	Data    OauthClientAccessTokenResData `json:"data"`
	Message string                        `json:"message"`
}

// OauthClientAccessToken 生成client_token
func (m *Manager) OauthClientAccessToken() (res OauthClientAccessTokenRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?grant_type=client_credential&client_key=%s&client_secret=%s", conf.API_OAUTH_CLIENT_ACCESS_TOKEN, m.Credentials.ClientKey, m.Credentials.ClientSecret), nil, nil)
	return res, err
}

// OauthRefreshTokenReq 刷新access_token请求
type OauthRefreshTokenReq struct {
	RefreshToken string // 填写通过access_token获取到的refresh_token参数
}

// OauthRefreshTokenResData 刷新access_token
type OauthRefreshTokenResData struct {
	AccessToken  string `json:"access_token"`  // 接口调用凭证
	Scope        string `json:"scope"`         // 用户授权的作用域
	ExpiresIn    uint64 `json:"expires_in"`    // 过期时间，单位（秒）
	OpenId       string `json:"open_id"`       // 当前应用下，授权用户唯一标识
	RefreshToken string `json:"refresh_token"` // 用户刷新
	DYError
}

// OauthRefreshTokenRes 刷新access_token
type OauthRefreshTokenRes struct {
	Data    OauthRefreshTokenResData `json:"data"`
	Message string                   `json:"message"`
}

// OauthRefreshToken 刷新access_token
func (m *Manager) OauthRefreshToken(req OauthRefreshTokenReq) (res OauthRefreshTokenRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?grant_type=refresh_token&client_key=%s&refresh_token=%s", conf.API_OAUTH_REFRESH_TOKEN, m.Credentials.ClientKey, req.RefreshToken), nil, nil)
	return res, err
}

// OauthRenewRefreshTokenReq 刷新refresh_token请求
type OauthRenewRefreshTokenReq struct {
	RefreshToken string // 填写通过access_token获取到的refresh_token参数
}

// OauthRenewRefreshTokenResData 刷新refresh_token
type OauthRenewRefreshTokenResData struct {
	ExpiresIn    uint64 `json:"expires_in"`    // 过期时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新
	DYError
}

// OauthRenewRefreshTokenRes 刷新refresh_token
type OauthRenewRefreshTokenRes struct {
	Data    OauthRenewRefreshTokenResData `json:"data"`
	Message string                        `json:"message"`
}

// OauthRenewRefreshToken 刷新refresh_token
func (m *Manager) OauthRenewRefreshToken(req OauthRenewRefreshTokenReq) (res OauthRenewRefreshTokenRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?client_key=%s&refresh_token=%s", conf.API_OAUTH_RENEW_REFRESH_TOKEN, m.Credentials.ClientKey, req.RefreshToken), nil, nil)
	return res, err
}

// OauthUserinfoReq 用户信息请求
type OauthUserinfoReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// OauthUserinfoResData 用户信息
type OauthUserinfoResData struct {
	Nickname      string `json:"nickname"`         // 昵称
	Province      string `json:"province"`         // 省
	Avatar        string `json:"avatar"`           // 头像
	City          string `json:"city"`             // 城市
	Country       string `json:"country"`          // 国家
	EAccountRole  string `json:"e_account_role"`   // 类型: * `EAccountM` - 普通企业号 * `EAccountS` - 认证企业号 * `EAccountK` - 品牌企业号
	Gender        int64  `json:"gender"`           // 性别: * `0` - 未知 * `1` - 男性 * `2` - 女性
	OpenId        string `json:"open_id"`          // 用户在当前应用的唯一标识
	UnionId       string `json:"union_id"`         // 用户在当前开发者账号下的唯一标识（未绑定开发者账号没有该字段）
	EncryptMobile string `json:"encrypt_mobile"`   // 手机号加密字符串
	Mobile        string `json:"mobile,omitempty"` // 手机号
	DYError
}

// OauthUserinfoRes 用户信息
type OauthUserinfoRes struct {
	Data OauthUserinfoResData `json:"data"`
}

// OauthUserinfo 获取用户信息
func (m *Manager) OauthUserinfo(req OauthUserinfoReq) (res *OauthUserinfoRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?open_id=%s&access_token=%s", conf.API_OAUTH_USERINFO, req.OpenId, req.AccessToken), nil, nil)
	if res.Data.EncryptMobile != "" {
		mobile, err := m.DecryptMobile(res.Data.EncryptMobile)
		if err != nil {
			return nil, err
		}
		res.Data.Mobile = mobile
	}
	return res, err
}

// DecryptMobile 解密用户手机号
func (m *Manager) DecryptMobile(encryptMobile string) (string, error) {
	oriData, err := Base64Decode(encryptMobile)
	if err != nil {
		return "", err
	}
	key := []byte(m.Credentials.ClientSecret)
	iv := key[:16]
	ret, err := AesDecrypt(oriData, key, iv)
	return string(ret), err
}
