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

// Env 实现Environment接口
type Env struct {
	value string
}

func (e *Env) Value() string {
	return e.value
}

func (e *Env) IsDev() bool {
	return e.value == "dev"
}

func (e *Env) IsFat() bool {
	return e.value == "fat"
}

func (e *Env) IsUat() bool {
	return e.value == "uat"
}

func (e *Env) IsPro() bool {
	return e.value == "pro"
}

// Active 方法获取当前配置的env环境
func Active() Environment {
	return active
}

func init() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
	flag.Parse()

	value := strings.ToLower(strings.TrimSpace(*env))
	envs := []string{"dev", "fat", "uat", "pro"}
	if !utils.Contain(value, envs) {
		//如果输入env变量不符合,那么启用默认fat环境
		active = &Env{value: "fat"}
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
}
