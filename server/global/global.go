package global

import (
	"go.uber.org/zap"
	//"gin-vue-admin/config"
	//"github.com/go-redis/redis"
	//"github.com/spf13/viper"
	"galaxy-golang/server/config"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	//GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	//GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
)
