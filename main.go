package main

import (
	"flag"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/internal/model"
	"goproject/main/ginweb/internal/routers"
	"goproject/main/ginweb/pkg/logger"
	"goproject/main/ginweb/pkg/setting"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	setUpFlag()
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

	//初始化链路跟踪
	// err = setupTracer()
	// if err != nil {
	// 	log.Fatalf("init.setupTracer err: %v", err)
	// }
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
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
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
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

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDbEngine() error {
	db, err := model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	a := gormadapter.NewAdapterByDB(db)
	global.DbEngine = db
	global.Casbin = casbin.NewEnforcer(global.AppSetting.RbacConfigPath, a)
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

// func setupTracer() error {
// 	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-server", "127.0.0.1:8888")
// 	if err != nil {
// 		return err
// 	}
// 	global.Trace = jaegerTracer
// 	return nil
// }

func setUpFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "./configs", "指定要要使用的配置文件路径")
	flag.Parse()
	return nil
}
