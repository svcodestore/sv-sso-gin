package initialize

import (
	"github.com/svcodestore/sv-sso-gin/service/system"
)

func WebsocketManagerInit() {
	system.WsClientMgr = system.NewWebsocketClientManager()
	go system.WsClientMgr.Start()
}
