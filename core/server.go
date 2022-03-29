package core

import (
	"fmt"
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/initialize"
	"github.com/svcodestore/sv-sso-gin/model"
	"log"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	global.CONFIGURATOR = initialize.InitConfigurator()
	global.LOGGER = initialize.Zap()
	global.DB = initialize.Gorm()
	initialize.DBList()
	if global.DB != nil {
		global.UserMgr = model.UsersMgr(global.DB)
		global.OrganizationMgr = model.OrganizationsMgr(global.DB)

		db, err := global.DB.DB()
		if err != nil {
			log.Panicln(err)
		}
		defer db.Close()
	}
	routers := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, routers)

	global.LOGGER.Error(s.ListenAndServe().Error())
}
