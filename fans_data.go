package douyingo

import (
	"context"

	"github.com/zhangshuai/douyin-go/conf"
)

// FansDataReq 用户粉丝数据请求
type FansDataReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// FansDataActiveDaysDistributions 用户粉丝数据
type FansDataActiveDaysDistributions struct {
	Item  string `json:"item"`  // 分布的种类
	Value int64  `json:"value"` // 分布的数值
}

// FansDataAgeDistributions 用户粉丝数据
type FansDataAgeDistributions struct {
	Item  string `json:"item"`  // 分布的种类
	Value int64  `json:"value"` // 分布的数值
}

// FansDataDeviceDistributions 用户粉丝数据
type FansDataDeviceDistributions struct {
	Item  string `json:"item"`  // 分布的种类
	Value int64  `json:"value"` // 分布的数值
}

// FansDataFlowContributions 用户粉丝数据
type FansDataFlowContributions struct {
	AllSum  int64  `json:"all_sum"`  // 总流量贡献
	FansSum int64  `json:"fans_sum"` // 粉丝流量贡献
	Flow    string `json:"flow"`     // 总流量贡献
}

// FansDataGenderDistributions 用户粉丝数据
type FansDataGenderDistributions struct {
	Item  string `json:"item"`  // 分布的种类
	Value int64  `json:"value"` // 分布的数值
}

// FansDataGeographicalDistributions 用户粉丝数据
type FansDataGeographicalDistributions struct {
	Item  string `json:"item"`  // 分布的种类
	Value int64  `json:"value"` // 分布的数值
}

// FansDataInterestDistributions 用户粉丝数据
type FansDataInterestDistributions struct {
	Item  string `json:"item"`  // 分布的种类
	Value int64  `json:"value"` // 分布的数值
}

// FansDataDetails 用户粉丝数据
type FansDataDetails struct {
	ActiveDays   []FansDataActiveDaysDistributions   `json:"active_days_distributions,omitempty"`  // 粉丝活跃天数分布 item: ["0-4","5~8","9~12","13~16","17~20","20+"]
	Age          []FansDataAgeDistributions          `json:"age_distributions,omitempty"`          // 粉丝年龄分布 item: ["1-23", "24-30", "31-40", "41-50", "50-"]
	Device       []FansDataDeviceDistributions       `json:"device_distributions,omitempty"`       // 粉丝设备分布 item: ["苹果","华为","三星","小米"...]
	Flow         []FansDataFlowContributions         `json:"flow_contributions,omitempty"`         // 粉丝流量贡献 flow: ["vv","like_cnt","comment_cnt","share_video_cnt"]
	Gender       []FansDataGenderDistributions       `json:"gender_distributions,omitempty"`       // 粉丝性别分布 item: ["1","2"] (男:1,女:2)
	Geographical []FansDataGeographicalDistributions `json:"geographical_distributions,omitempty"` // 粉丝地域分布 item: ["北京","福建","香港"...]
	Interest     []FansDataInterestDistributions     `json:"interest_distributions,omitempty"`     // 粉丝兴趣分布 item: ["生活"","美食","旅行"...]
	AllFansNum   int64                               `json:"all_fans_num,omitempty"`               // 所有粉丝的数量
}

// FansDataList 用户粉丝数据
type FansDataList struct {
	Details FansDataDetails `json:"fans_data,omitempty"`
	DYError
}

// FansDataRes 用户粉丝数据
type FansDataRes struct {
	Data  FansDataList `json:"data"`
	Extra DYExtra      `json:"extra"`
}

// FansData 获取用户粉丝数据(用户首次授权应用后，需要间隔2天才会产生全部的数据；并只提供粉丝大于100的用户数据。)
func (m *Manager) FansData(req FansDataReq) (res FansDataRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s", conf.API_FANS_DATA, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}

// DataExternalFansSourceReq 用户粉丝来源分布请求
type DataExternalFansSourceReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// DataExternalFansSourceItem 用户粉丝来源分布
type DataExternalFansSourceItem struct {
	Source  string `json:"source"`  // 粉丝来源
	Percent string `json:"percent"` // 来源占比，格式:XX.XX% 精确小数点后2位
}

// DataExternalFansSourceData 用户粉丝来源分布
type DataExternalFansSourceData struct {
	List []DataExternalFansSourceItem `json:"list,omitempty"`
	DYError
}

// DataExternalFansSourceRes 用户粉丝来源分布
type DataExternalFansSourceRes struct {
	Data  DataExternalFansSourceData `json:"data"`
	Extra DYExtra                    `json:"extra"`
}

// DataExternalFansSource 获取用户粉丝来源分布
func (m *Manager) DataExternalFansSource(req DataExternalFansSourceReq) (res DataExternalFansSourceRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s", conf.API_DATA_EXTERNAL_FANS_SOURCE, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}

// DataExternalFansFavouriteReq 用户粉丝喜好请求
type DataExternalFansFavouriteReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// DataExternalFansFavouriteItem 用户粉丝喜好
type DataExternalFansFavouriteItem struct {
	Rank     int32  `json:"rank"`      // 排名
	Keyword  string `json:"keyword"`   // 关键词
	HotValue int64  `json:"hot_value"` // 热度指数
}

// DataExternalFansFavouriteData 用户粉丝喜好
type DataExternalFansFavouriteData struct {
	List []DataExternalFansFavouriteItem `json:"list,omitempty"`
	DYError
}

// DataExternalFansFavouriteRes 用户粉丝喜好
type DataExternalFansFavouriteRes struct {
	Data  DataExternalFansFavouriteData `json:"data"`
	Extra DYExtra                       `json:"extra"`
}

// DataExternalFansFavourite 获取用户粉丝喜好
func (m *Manager) DataExternalFansFavourite(req DataExternalFansFavouriteReq) (res DataExternalFansFavouriteRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s", conf.API_DATA_EXTERNAL_FANS_FAVOURITE, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}

// DataExternalFansCommentReq 用户粉丝热评
type DataExternalFansCommentReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

// DataExternalFansCommentItem 用户粉丝热评
type DataExternalFansCommentItem struct {
	Rank     int32  `json:"rank"`      // 排名
	Keyword  string `json:"keyword"`   // 关键词
	HotValue int64  `json:"hot_value"` // 热度指数
}

// DataExternalFansCommentData 用户粉丝热评
type DataExternalFansCommentData struct {
	List []DataExternalFansCommentItem `json:"list,omitempty"`
	DYError
}

// DataExternalFansCommentRes 用户粉丝热评
type DataExternalFansCommentRes struct {
	Data  DataExternalFansCommentData `json:"data"`
	Extra DYExtra                     `json:"extra"`
}

// DataExternalFansComment 获取用户粉丝热评
func (m *Manager) DataExternalFansComment(req DataExternalFansCommentReq) (res DataExternalFansCommentRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s", conf.API_DATA_EXTERNAL_FANS_COMMENT, req.AccessToken, req.OpenId), nil, nil)
	return res, err
}
