package douyingo

import (
	"context"

	"github.com/guaidashu/douyin-go/conf"
)

// RoleCheckReq 用户经营身份管理请求
type RoleCheckReq struct {
	AccessToken   string   `json:"_"`                        // 调用/oauth/client_token/生成的 token，此 token 不需要用户授权。
	OpenId        string   `json:"open_id,omitempty"`        // 用户open_id, 通过/oauth/access_token/获取，用户唯一标志。与douyin_shortId字段必须传一个，如果都传会以open_id为第一优先级进行查询。
	DouyinShortId string   `json:"douyin_shortId,omitempty"` // 用户抖音号。与open_id字段必须传一个，如果都传会以open_id为第一优先级进行查询。
	RoleLabels    []string `json:"role_labels,omitempty"`    // 多个身份信息枚举数组，可以根据需求传以下字段列表：COMPANY_BAND：企业号品牌号 AUTH_COMPANY：认证企业号 STAFF：员工号 OPEN_BRAND：开平品牌号 OPEN_STAFF：开平员工号 OPEN_PARTNER：开平合作号
}

// RoleCheckData 用户经营身份管理响应数据
type RoleCheckData struct {
	MatchResult bool            `json:"match_result,omitempty"` // 匹配状态
	FilterRole  map[string]bool `json:"filter_role,omitempty"`  // 返回查询集合中，没有身份的信息
}

// RoleCheckRes 用户经营身份管理响应
type RoleCheckRes struct {
	Data RoleCheckData `json:"data"`
	DYExtraV1
}

// RoleCheck 获取用户经营身份管理
func (m *Manager) RoleCheck(req RoleCheckReq) (res RoleCheckRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s", conf.API_ROLE_CHECK, req.AccessToken), nil, req)
	return res, err
}
