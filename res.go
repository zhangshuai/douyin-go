package douyingo

// DYExtra 响应结构体
type DYExtra struct {
	LogId         string `json:"logid"`                     // 标识请求的唯一id
	Now           uint64 `json:"now"`                       // 毫秒级时间戳
	SubDescrition string `json:"sub_description,omitempty"` // 子错误码描述
	SubErrorCode  int64  `json:"sub_error_code,omitempty"`  // 子错误码
	DYError
}

// DYExtraV1 V1响应结构体
type DYExtraV1 struct {
	LogId  string `json:"log_id"`            // 标识请求的唯一id
	ErrNo  int64  `json:"err_no,omitempty"`  // 错误码
	ErrMsg string `json:"err_msg,omitempty"` // 错误信息
}
