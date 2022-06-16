package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

/*
配置信息进行全局变量的声明，便于在接下来的步骤将其关联起来，并且提供给应用程序内部调用
另外全局变量的初始化，是会随着应用程序的不断演进不断改变的，因此并不是一成不变，也就是这里展示的并不一定是最终的结果。
*/

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS      //配置信息的全局变量对象
	DatabaseSetting *setting.DatabaseSettingS //数据库的全局变量对象
	Logger          *logger.Logger            //日志的全局变量对象
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	ZapLogger       *setting.ZapLoggers
)
