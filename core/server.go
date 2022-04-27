package core

import (
	"fmt"
	"github.com/svcodestore/sv-sso-gin/rpc"
	"log"

	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/initialize"
	"github.com/svcodestore/sv-sso-gin/model"
)

type server interface {
	ListenAndServe() error
}

func commonInit() {
	global.CONFIGURATOR = initialize.InitConfigurator()
	global.LOGGER = initialize.Zap()
	global.DB = initialize.Gorm()
	initialize.DBList()
	if global.DB != nil {
		global.UserMgr = model.UsersMgr(global.DB)
		global.OrganizationMgr = model.OrganizationsMgr(global.DB)
		global.ApplicationMgr = model.ApplicationsMgr(global.DB)
		global.ApplicationUserMgr = model.ApplicationUserMgr(global.DB)
		global.OrganizationApplicationMgr = model.OrganizationApplicationMgr(global.DB)

		db, err := global.DB.DB()
		if err != nil {
			log.Panicln(err)
		}
		defer db.Close()
	}

	initialize.Redis()
}

func RunServer() {
	commonInit()

	routers := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, routers)

	global.LOGGER.Error(s.ListenAndServe().Error())

}

func RunRpcServer()  {
	commonInit()
	initialize.InitGrpc(rpc.RegisterServer)
}