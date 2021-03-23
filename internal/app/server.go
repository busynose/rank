package app

import (
	"rank/configs"
	"net/http"
	"time"

	"gitlab.com/makeblock-go/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	maxAgeInMinutes  = 10
	timeoutInSeconds = 30
)

//RunServer run server with port
func RunServer(port string) {
	setGinMode()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION"},
		AllowHeaders:     []string{"utoken,x-auth-token,x-request-id,Content-Type,Accept,Origin,Access-Control-Allow-Origin", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           maxAgeInMinutes * time.Minute,
	}))
	registerRouters(router)
	startServer(router, port)
}

func setGinMode() {
	env := configs.Env.ProjectEnv
	if env == configs.Dev {
		gin.SetMode(gin.DebugMode)
	} else if env == configs.Test {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func startServer(router *gin.Engine, port string) {
	// Listen and serve on 0.0.0.0:8080
	// router.Run(":80") 这样写就可以了，下面所有代码（go1.8+）是为了优雅处理重启等动作。可根据实际情况选择。
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  timeoutInSeconds * time.Second,
		WriteTimeout: timeoutInSeconds * time.Second,
	}

	go func() {
		log.Println("Start Http Server ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.FatalE("Failed to serve: ", err)
		}
	}()
}
