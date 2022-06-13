package system

import (
	"github.com/gorilla/websocket"
	"log"
	"runtime/debug"
)

type WebsocketClient struct {
	Conn     *websocket.Conn
	ToBeSent chan []byte
	UserId   string
}

func NewWebsocketClient(conn *websocket.Conn) (client *WebsocketClient) {
	client = &WebsocketClient{
		Conn:     conn,
		ToBeSent: make(chan []byte, 100),
	}

	return
}

func (c *WebsocketClient) Read() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("read stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		close(c.ToBeSent)
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			isUnset := oauthService.UnsetUserOnline(c.UserId)
			if isUnset {
				log.Printf("user[%s] deactive successful", c.UserId)
				users := oauthService.AllOnlineUser()
				b, _ := json.Marshal(users)
				rtn := []byte("onlineUsers:")
				rtn = append(rtn, b...)
				WsClientMgr.BroadcastMessage <- rtn
			} else {
				log.Printf("user[%s] deactive fail", c.UserId)
			}
			break
		}
		log.Printf("recv user[%s]: %s", c.UserId, message)

		if string(message) == "ping" {
			c.SendMsg([]byte("pong"))
		}
	}
}

func (c *WebsocketClient) Write() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("write stop", string(debug.Stack()), r)
		}
	}()

	for {
		select {
		case message, ok := <-c.ToBeSent:
			if !ok {
				log.Println("client 发送数据 关闭连接", ok)

				return
			}

			err := c.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func (c *WebsocketClient) SendMsg(msg []byte) {
	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.ToBeSent <- msg
}
