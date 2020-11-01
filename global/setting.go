package global

import (
	"github.com/jinzhu/gorm"
	"github.com/vinsec/blog-service/pkg/logger"
	"github.com/vinsec/blog-service/pkg/setting"
)

//Instantiate global objects that save configs
var (
	ServerSetting		*setting.ServerSettingS
	AppSetting			*setting.AppSettingS
	DatabaseSetting		*setting.DataBaseSettingS
	DBEngine			*gorm.DB
	Logger				*logger.Logger
)


