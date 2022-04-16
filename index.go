package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/model"
	"goproject/main/ginweb/internal/routers"
	"goproject/main/ginweb/pkg/logger"
	"goproject/main/ginweb/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	//初始化配置文件
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//初始化数据库
	DbErr := setupDbEngine()
	if DbErr != nil {
		log.Fatalf("init.NewDBEngine err: %v", err)
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	router := routers.NewRoutes()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDbEngine() error {
	db, err := model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DbEngine = db
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
