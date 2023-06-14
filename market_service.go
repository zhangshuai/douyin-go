package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// MarketServicePurchaseListReq 查询用户的服务购买信息请求
type MarketServicePurchaseListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户的唯一标志
	AccessToken string // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	ServiceId   string // 服务id，服务的唯一标识，可在发布“平台应用类”功能中【应用配置页】和【服务详情页】中获取
	IsTestEnv   bool   // 是否为测试环境：true为测试环境，在服务未发布前或需要测试时使用测试环境参数; false为正式环境，需要传入的服务已发布
}

// MarketServicePurchaseListResult 用户的服务购买信息
type MarketServicePurchaseListResult struct {
	ServiceId string `json:"service_id"` // 服务id，服务的唯一标识
}

// MarketServicePurchaseListData 用户的服务购买信息
type MarketServicePurchaseListData struct {
	List []MarketServicePurchaseListResult `json:"purchase_info_list"` // 购买信息列表
	DYError
}

// MarketServicePurchaseListRes 用户的服务购买信息
type MarketServicePurchaseListRes struct {
	Data  MarketServicePurchaseListData `json:"data"`
	Extra DYExtra                       `json:"extra"`
}

// MarketServicePurchaseList 查询用户的服务购买信息
func (m *Manager) MarketServicePurchaseList(req MarketServicePurchaseListReq) (res MarketServicePurchaseListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&service_id=%s&is_test_env=%t", conf.API_VIDEO_SEARCH, req.AccessToken, req.OpenId, req.ServiceId, req.IsTestEnv), nil, nil)
	return res, err
}
