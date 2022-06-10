package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/service/system"
	"log"
	"net/http"
	"strings"
)

func UserActivation(c *gin.Context) {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	userId := c.Param("id")
	isUserLogin := oauthService.IsUserLogin(userId)
	log.Printf("is user[%s] login: %v", userId, isUserLogin)
	if !isUserLogin {
		conn.Close()
		return
	}
	isUserOnline := oauthService.IsUserOnline(userId)
	if !isUserOnline {
		user, _ := userService.UserWithId(userId)
		oauthService.SetUserOnline(userId, user.Name)
	}

	client := system.NewWebsocketClient(conn)
	client.UserId = userId

	go client.Read()
	go client.Write()

	log.Println(service.ServiceGroup.WebsocketClientManagerService.ManagerInfo())
	service.ServiceGroup.WebsocketClientManagerService.Connect <- client
	log.Println(service.ServiceGroup.WebsocketClientManagerService.ManagerInfo())
	users := oauthService.AllOnlineUser()
	log.Println(users)
	service.ServiceGroup.WebsocketClientManagerService.BroadcastMessage <- []byte(strings.Join(users, ","))
	log.Println(service.ServiceGroup.WebsocketClientManagerService.ManagerInfo())
}
