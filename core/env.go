package core

import (
	"flag"
	"fmt"
	"strings"

	"github.com/anthonyzero/go-quick-api/utils"
)

var (
	active Environment
)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsFat() bool
	IsUat() bool
	IsPro() bool
}

// env 实现Environment接口
type env struct {
	value string
}

func (e *env) Value() string {
	return e.value
}

func (e *env) IsDev() bool {
	return e.value == "dev"
}

func (e *env) IsFat() bool {
	return e.value == "fat"
}

func (e *env) IsUat() bool {
	return e.value == "uat"
}

func (e *env) IsPro() bool {
	return e.value == "pro"
}

// Active 方法获取当前配置的env环境
func Active() Environment {
	return active
}

func init() {
	envstr := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
	flag.Parse()

	value := strings.ToLower(strings.TrimSpace(*envstr))
	envs := []string{"dev", "fat", "uat", "pro"}
	if !utils.Contain(value, envs) {
		//如果输入env变量不符合,那么启用默认fat环境
		active = &env{value: "fat"}
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
}
