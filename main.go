package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"blog-service/pkg/tracer"
	"blog-service/pkg/zapLogger"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//初始化数据库配置
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
    //初始化Uber zap
	err=setZapLogger()
	if err != nil {
		log.Fatalf("init.setZapLogger err: %v", err)
	}

	//
	err=setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:{0}",err)
	}
}

// @title 博客系统
//@version 1.0
//@description Go 语言编程之旅：一起用 Go 做项目
//@termsOfService  https://github.com/go-programming-tour-book
func main() {


	gin.SetMode(global.ServerSetting.RunMode)
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	router := routers.NewRouter()
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
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("ZapLogger", &global.ZapLogger)
	if err != nil {
		return err
	}


	global.ServerSetting.HttpPort = "8080"
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second


	return nil
}

//初始化全局变量的数据路连接池
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

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

func setZapLogger()  error{
	if err := zapLogger.InitLogger(global.ZapLogger); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return err
	}
	return nil
}

func setupTracer() error {
	jaegerTracr,_,err:=tracer.NewJaegerTracer("blog-service","127.0.0.1:6831")
	if err!=nil{
		return err
	}
	global.Tracer=jaegerTracr
	return nil
}




//最好的投资是投资自己 当你某项能力学到足以打败大多数人的时候 牛马人的命运自然迎刃而解
//通过学习获得杠杆
