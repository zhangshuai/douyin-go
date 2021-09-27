package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// PoiSearchKeywordReq POI信息请求
type PoiSearchKeywordReq struct {
	AccessToken string // 调用/oauth/client_token/生成的token，此token不需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	Keyword     string // 查询关键字，例如美食
	City        string // 城市
}

// Poi POI信息
type Poi struct {
	CountryCode string `json:"country_code"` // 国家编码
	PoiId       string `json:"poi_id"`       // 唯一ID
	District    string `json:"district"`     // 区域名称
	Location    string `json:"location"`     // 经纬度，格式：X,Y
	PoiName     string `json:"poi_name"`     // 名称
	Province    string `json:"province"`     // 省份
	Address     string `json:"address"`      // 地址
	City        string `json:"city"`         // 城市
	CityCode    string `json:"city_code"`    // 城市编码
	Country     string `json:"country"`      // 国家
}

// PoiSearchKeywordData POI信息
type PoiSearchKeywordData struct {
	Pois    []Poi `json:"pois"`
	Cursor  int64 `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool  `json:"has_more"` // 更多数据
	DYError
}

// PoiSearchKeywordRes POI信息
type PoiSearchKeywordRes struct {
	Data  PoiSearchKeywordData `json:"data"`
	Extra DYExtra              `json:"extra"`
}

// PoiSearchKeyword 查询POI信息
func (m *Manager) PoiSearchKeyword(req PoiSearchKeywordReq) (res PoiSearchKeywordRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&cursor=%d&count=%d&keyword=%s&city=%s", conf.API_POI_SEARCH_KEYWORD, req.AccessToken, req.Cursor, req.Count, req.Keyword, req.City), nil, nil)
	return res, err
}
