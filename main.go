package main

import (
	"jassue-gin/bootstrap"
	"jassue-gin/global"
)

func main() {
	bootstrap.InitializeConfig()

	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	global.App.DB = bootstrap.InitializeDB()

	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	bootstrap.InitializeValidator()

	bootstrap.RunServer()
}
