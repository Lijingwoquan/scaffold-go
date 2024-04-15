package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"scaffold/dao/mysql"
	"scaffold/logger"
	"scaffold/pkg/snowflake"
	ticker "scaffold/pkg/tickerTask"
	"scaffold/routers"
	"scaffold/setting"
)

func main() {
	//1.加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Println("init setting failed!")
	}

	//2.初始化日志
	if err := logger.Init(viper.GetString("app.mode")); err != nil {
		return
	}
	defer func() {
		if err := zap.L().Sync(); err != nil {
			fmt.Printf(" zap.L().Sync() failed,err:%v", err)
		}
	}()

	//3.初始化mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed! err:%v", err)
		return
	}

	//初始化雪花算法
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("snowflake init failed,err:%v", err)
		return
	}

	//初始计时器
	ticker.Init()

	//5.注册路由
	r := routers.SetupRouter(viper.GetString("app.mode"))
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("run server failed,err:%v", err)
		return
	}
}
