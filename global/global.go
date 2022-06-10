package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/svcodestore/sv-sso-gin/config"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	DBList             map[string]*gorm.DB
	REDIS              *redis.Client
	RpcServer          *grpc.Server
	CONFIGURATOR       config.Configurator
	CONFIG             config.Config
	LOGGER             *zap.Logger
	ConcurrencyControl = &singleflight.Group{}
)
