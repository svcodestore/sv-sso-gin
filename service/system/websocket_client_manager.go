package system

import (
	"sync"
)

var WsClientMgr *WebsocketClientManager

type WebsocketClientManager struct {
	Clients          map[string]*WebsocketClient
	Lock             sync.RWMutex
	Connect          chan *WebsocketClient
	Disconnect       chan *WebsocketClient
	BroadcastMessage chan []byte
}

func NewWebsocketClientManager() (manager *WebsocketClientManager) {
	manager = &WebsocketClientManager{
		Clients:          make(map[string]*WebsocketClient),
		Connect:          make(chan *WebsocketClient, 2<<10),
		Disconnect:       make(chan *WebsocketClient, 2<<10),
		BroadcastMessage: make(chan []byte, 1000),
	}

	return
}

func (m *WebsocketClientManager) ClientWithUserId(userId string) (client *WebsocketClient) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	client = m.Clients[userId]

	return
}

func (m *WebsocketClientManager) ClientCount() (count int) {
	count = len(m.Clients)

	return
}

func (m *WebsocketClientManager) DisconnectClient(userId string) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	if _, ok := m.Clients[userId]; ok {
		delete(m.Clients, userId)
	}
}

func (m *WebsocketClientManager) ConnectClient(userId string, c *WebsocketClient) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	m.Clients[userId] = c
}

func (m *WebsocketClientManager) ManagerInfo() (managerInfo map[string]interface{}) {
	managerInfo = make(map[string]interface{})

	managerInfo["clientsCount"] = m.ClientCount()
	managerInfo["ConnectionsCount"] = len(m.Connect)
	managerInfo["DisconnectionsCount"] = len(m.Disconnect)
	managerInfo["BroadcastCount"] = len(m.BroadcastMessage)

	return
}

func (m *WebsocketClientManager) Start() {
	for {
		select {
		case conn := <-m.Connect:
			m.ConnectClient(conn.UserId, conn)
		case conn := <-m.Disconnect:
			m.DisconnectClient(conn.UserId)
		case message := <-m.BroadcastMessage:
			users := oauthService.AllOnlineUser()
			for id, _ := range users {
				if client, ok := m.Clients[id]; ok {
					select {
					case client.ToBeSent <- message:
					default:
						close(client.ToBeSent)
					}
				}
			}
		}
	}
}
