package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/svcodestore/sv-sso-gin/service/system"
	"github.com/svcodestore/sv-sso-gin/utils"
	"log"
	"net/http"
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
	token := c.Query("token")
	if token == "" {
		return
	}
	j := utils.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil || claims.UserId != userId {
		return
	}
	isUserLogin := oauthService.IsUserLogin(userId)
	log.Printf("is user[%s] login: %v", userId, isUserLogin)
	if !isUserLogin {
		return
	}
	isUserOnline := oauthService.IsUserOnline(userId)
	if !isUserOnline {
		user, _ := userService.UserWithId(userId)
		oauthService.SetUserOnline(userId, user.Name)
	}

	client := system.NewWebsocketClient(conn)
	client.UserId = userId

	system.WsClientMgr.Connect <- client
	users := oauthService.AllOnlineUser()
	b, _ := json.Marshal(users)
	rtn := []byte("onlineUsers:")
	rtn = append(rtn, b...)
	system.WsClientMgr.BroadcastMessage <- rtn

	 go client.Read()
	 client.Write()

}
