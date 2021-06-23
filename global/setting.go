package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
	"github.com/opentracing/opentracing-go"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting *setting.JWTSettingS
	EmailSetting *setting.EmailSettingS
	Tracer opentracing.Tracer

	Logger *logger.Logger
)