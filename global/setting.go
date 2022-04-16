package global

import (
	"goproject/main/ginweb/pkg/logger"
	"goproject/main/ginweb/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
