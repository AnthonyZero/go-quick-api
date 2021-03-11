package core

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

// Env 方法获取当前配置的env环境
func Env() Environment {
	return active
}

//Setup 设置环境
func EnvSetup(value string) {
	active = &env{value: value}
}
