package main

import (
	"os"
	"os/signal"
	"rank/configs"
	"rank/internal/app"
	"rank/migrate"

	"gitlab.com/makeblock-go/log"
	"gitlab.com/makeblock-go/mysql"
	"gitlab.com/makeblock-go/redis"
)

func main() {
	isProduction := configs.Env.ProjectEnv == configs.Prod
	log.SetUp(isProduction,
		log.Any("serverName", "rank"))
	defer log.Sync()

	cnf := mysql.NewConfig(
		configs.Env.Mysql.User,
		configs.Env.Mysql.Pwd,
		configs.Env.Mysql.Host,
		configs.Env.Mysql.Port,
		configs.Env.Mysql.DBName,
		configs.Env.Mysql.Charset,
		!isProduction)
	mysql.Register(cnf)
	defer mysql.Close()

	redis.SetUp(
		configs.Env.Redis.Host,
		configs.Env.Redis.Port,
		configs.Env.Redis.Pwd)
	defer redis.Close()

	migrate.InitModel()
	migrate.InitData()

	app.RunServer(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
