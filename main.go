package main

import (
	"github.com/vinsec/blog-service/global"
	"github.com/vinsec/blog-service/internal/model"
	"github.com/vinsec/blog-service/internal/routers"
	"github.com/vinsec/blog-service/pkg/logger"
	"github.com/vinsec/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init(){
	err := setupSetting()
	if err != nil{
		log.Fatalf("init.setupSetting err: %v",err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v",err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v",err)
	}
}

//unmarshal config sections from FILE to the global objects
func setupSetting()error{
	settings,err := setting.NewSetting()
	if err != nil{
		return err
	}
	err = settings.ReadSection("Server",&global.ServerSetting)
	if err != nil{
		return err
	}
	err = settings.ReadSection("App",&global.AppSetting)
	if err != nil{
		return err
	}
	err = settings.ReadSection("Database",&global.DatabaseSetting)
	if err != nil{
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func setupLogger()error{
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)

	return nil
}

//Use "global.DatabaseSetting" to initiate the "global.DBEngine"
func setupDBEngine()error{
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func main(){
	router := routers.NewRouter()
	server := &http.Server{
		Addr:				":8080",
		Handler: 			router,
		ReadTimeout: 		10 * time.Second,
		WriteTimeout: 		10 * time.Second,
		MaxHeaderBytes: 	1 << 20,
	}
	server.ListenAndServe()
}
