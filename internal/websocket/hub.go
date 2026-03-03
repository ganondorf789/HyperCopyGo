package websocket

import (
	"sync"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

// WsMessage WebSocket 消息结构
type WsMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// Client 代表一个 WebSocket 连接
type Client struct {
	ws *ghttp.WebSocket
	mu sync.Mutex
}

// Write 向该客户端发送消息（并发安全）
func (c *Client) Write(msg WsMessage) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ws.WriteMessage(ghttp.WsMsgText, msgBytes)
}

// Hub 管理所有 WebSocket 连接，支持按 userId 推送
type Hub struct {
	mu      sync.RWMutex
	clients map[int64]map[*Client]struct{}
}

var defaultHub = &Hub{
	clients: make(map[int64]map[*Client]struct{}),
}

// GetHub 获取全局 Hub 实例
func GetHub() *Hub {
	return defaultHub
}

// Register 注册一个新的 WebSocket 连接
func (h *Hub) Register(userId int64, ws *ghttp.WebSocket) *Client {
	c := &Client{ws: ws}
	h.mu.Lock()
	if h.clients[userId] == nil {
		h.clients[userId] = make(map[*Client]struct{})
	}
	h.clients[userId][c] = struct{}{}
	h.mu.Unlock()
	return c
}

// Unregister 移除一个 WebSocket 连接
func (h *Hub) Unregister(userId int64, c *Client) {
	h.mu.Lock()
	if clients, ok := h.clients[userId]; ok {
		delete(clients, c)
		if len(clients) == 0 {
			delete(h.clients, userId)
		}
	}
	h.mu.Unlock()
}

// SendToUser 向指定用户的所有连接推送消息
func (h *Hub) SendToUser(userId int64, msg WsMessage) {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return
	}
	h.mu.RLock()
	clients := h.clients[userId]
	h.mu.RUnlock()

	for c := range clients {
		c.mu.Lock()
		_ = c.ws.WriteMessage(ghttp.WsMsgText, msgBytes)
		c.mu.Unlock()
	}
}

// Broadcast 向所有在线用户广播消息
func (h *Hub) Broadcast(msg WsMessage) {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return
	}
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, clients := range h.clients {
		for c := range clients {
			c.mu.Lock()
			_ = c.ws.WriteMessage(ghttp.WsMsgText, msgBytes)
			c.mu.Unlock()
		}
	}
}

// OnlineUserCount 返回当前在线用户数
func (h *Hub) OnlineUserCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}

// IsUserOnline 检查指定用户是否在线
func (h *Hub) IsUserOnline(userId int64) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	clients, ok := h.clients[userId]
	return ok && len(clients) > 0
}
