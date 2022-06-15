package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/svcodestore/sv-sso-gin/service/system"
	"github.com/svcodestore/sv-sso-gin/utils"
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

	ip4 := ""
	ip6 := ""
	clientIp := c.ClientIP()
	ips := strings.Split(clientIp, ":")
	ip4 = ips[len(ips)-1]
	device := ""
	clientId := claims.BaseClaims.ClientId
	application, _ := applicationService.ApplicationWithClientId(clientId)
	userLoginRecordService.UpsertUserLogin(userId, application.ID, ip4, ip6, device)

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
