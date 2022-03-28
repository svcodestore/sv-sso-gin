package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/svcodestore/sv-sso-gin/config"
	"github.com/svcodestore/sv-sso-gin/model"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	DBList             map[string]*gorm.DB
	REDIS              *redis.Client
	CONFIGURATOR       config.Configurator
	CONFIG             config.Config
	LOGGER             *zap.Logger
	ConcurrencyControl = &singleflight.Group{}
	UserMgr            model.UsersMgrType
)
