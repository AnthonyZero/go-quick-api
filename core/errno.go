package core

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// 实现编译期间检测接口是否实现。
var _ Error = (*err)(nil)

// Error 用来生成统一的 Response Body 的 JSON 数据。
type Error interface {
	// i 为了避免被其他包实现
	i()
	// WithErr 设置错误信息
	WithErr(err error) Error
	// GetCode 获取 业务Code
	GetCode() int
	// GetHttpCode 获取 HTTP Code
	GetHTTPCode() int
	// GetMsg 获取 Msg
	GetMsg() string
	// GetErr 获取错误信息
	GetErr() error
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	HTTPCode int    // HTTP Code
	Code     int    // 业务 Code
	Message  string // 描述信息
	Err      error  // 错误信息
}

// NewError 实例化一个错误码
func NewError(httpCode, code int, msg string) Error {
	return &err{
		HTTPCode: httpCode,
		Code:     code,
		Message:  msg,
	}
}

func (e *err) i() {}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHTTPCode() int {
	return e.HTTPCode
}

func (e *err) GetCode() int {
	return e.Code
}

func (e *err) GetMsg() string {
	return e.Message
}

func (e *err) GetErr() error {
	return e.Err
}

func (e *err) ToString() string {
	err := &struct {
		HTTPCode int    `json:"http_code"`
		Code     int    `json:"code"`
		Message  string `json:"message"`
	}{
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
