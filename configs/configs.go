package configs

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//config 项目配置文件信息
type config struct {
	//Base
	Base struct {
		Port    string `toml:"port"`
		AppName string `toml:"appName"`
		LogPath string `toml:"logPath"`
	} `toml:"base"`
	//Mysql配置
	MySQL struct {
		Default struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"default"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"mysql"`
	//Redis配置
	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`

	JWT struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"jwt"`

	Aes struct {
		Key string `toml:"key"`
		Iv  string `toml:"iv"`
	} `toml:"aes"`

	Rsa struct {
		Private string `toml:"private"`
		Public  string `toml:"public"`
	} `toml:"rsa"`

	Cmd struct {
		GenTables string `toml:"genTables"`
	} `toml:"cmd"`
}

//private
var configs = new(config)

//Get 获取配置文件 public
func Get() config {
	return *configs
}

//Init 配置初始化 active -> Env(dev fat uat pro)
func Init(active string) {
	viper.SetConfigName(active + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		// viper解析文件错误
		panic(err)
	}

	if err := viper.Unmarshal(configs); err != nil {
		panic(err)
	}

	// 监控配置文件变化
	watchConfig()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config file changed : %s", in.Name)
	})
}

//AppName 获取appname
func AppName() string {
	return configs.Base.AppName
}

//ServerPort 获取服务端口
func ServerPort() string {
	return configs.Base.Port
}
