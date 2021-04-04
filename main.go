package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anthonyzero/go-quick-api/configs"
	"github.com/anthonyzero/go-quick-api/core"
	"github.com/anthonyzero/go-quick-api/pkg/cache"
	"github.com/anthonyzero/go-quick-api/pkg/color"
	"github.com/anthonyzero/go-quick-api/pkg/db"
	"github.com/anthonyzero/go-quick-api/pkg/logger"
	"github.com/anthonyzero/go-quick-api/pkg/shutdown"
	"github.com/anthonyzero/go-quick-api/router"
	"github.com/anthonyzero/go-quick-api/utils"
	"github.com/gin-gonic/gin"
)

var (
	config = flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
)

func main() {
	flag.Parse()
	if *config == "" {
		core.EnvSetup("dev")
		log.Println(color.Yellow("Warning: '-env' cannot be found, The default 'dev' will be used."))
	} else {
		envs := []string{"dev", "fat", "uat", "pro"}
		if !utils.Contain(*config, envs) {
			log.Println(color.Red("Error: '-env' it is illegal"))
			flag.Usage()
			os.Exit(1)
		}
		core.EnvSetup(*config)
	}
	//配置文件加载
	configs.Init(core.Env().Value())
	//日志初始化
	logger.InitZapLogger(configs.Get().Base.LogPath, logger.ToLevel("info"))

	//初始化db
	if err := db.InitMysql(); err != nil {
		log.Println(color.Red(fmt.Sprintf("mysql init error: %s", err)))
		os.Exit(1)
	}

	//初始化redis
	if err := cache.InitRedis(); err != nil {
		log.Println(color.Red(fmt.Sprintf("redis init error: %s", err)))
		os.Exit(1)
	}

	// 初始化 HTTP 服务
	gin.SetMode(gin.DebugMode)
	engine := router.LoadRouter()
	server := &http.Server{
		Addr:    ":" + configs.ServerPort(),
		Handler: engine,
	}

	go func() {
		log.Println(color.Blue(fmt.Sprintf("[INFO] HttpServerRun:%s", configs.ServerPort())))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(color.Red(fmt.Sprintf("[ERROR] HttpServerRun:%s err: %v", configs.ServerPort(), err)))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Println(color.Red(fmt.Sprintf("server shutdown err %v", err)))
			} else {
				log.Println(color.Blue("http server shutdown success"))
			}
		},

		// 关闭 db
		func() {
			if err := db.Close(); err != nil {
				log.Println(color.Red(fmt.Sprintf("db close err %v", err)))
			} else {
				log.Println(color.Blue("db close success"))
			}
		},

		// 关闭 cache
		func() {
			if err := cache.Close(); err != nil {
				log.Println(color.Red(fmt.Sprintf("cache close err %v", err)))
			} else {
				log.Println(color.Blue("cache close success"))
			}
		},
	)
}
