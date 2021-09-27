package douyingo

import "fmt"

// DYError 错误结构体
type DYError struct {
	ErrorCode   int64  `json:"error_code,omitempty"`  // 错误码
	Description string `json:"description,omitempty"` // 错误码描述
}

func (e *DYError) Error() string {
	return fmt.Sprintf("%d: %s", e.ErrorCode, e.Description)
}

// NewError 新建错误结构体
func NewError(errorCode int64, description string) *DYError {
	return &DYError{
		ErrorCode:   errorCode,
		Description: description,
	}
}
