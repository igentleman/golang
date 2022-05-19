package global

import (
	"goproject/main/ginweb/pkg/logger"
	"goproject/main/ginweb/pkg/setting"

	"github.com/casbin/casbin"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	Casbin          *casbin.Enforcer
)
