package event

import (
	"sync"
)

// Hub 用于管理日志事件的发布和订阅
type Hub struct {
	subscribers map[string][]chan string
	mu          sync.RWMutex
}

// NewHub 创建一个新的 Hub 实例
func NewHub() *Hub {
	return &Hub{
		subscribers: make(map[string][]chan string),
	}
}

// Subscribe 订阅特定类型的事件
func (h *Hub) Subscribe(eventType string) chan string {
	ch := make(chan string)
	h.mu.Lock()
	h.subscribers[eventType] = append(h.subscribers[eventType], ch)
	h.mu.Unlock()
	return ch
}

// Publish 发布事件
func (h *Hub) Publish(eventType string, message string) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if subscribers, found := h.subscribers[eventType]; found {
		for _, ch := range subscribers {
			ch <- message
		}
	}
}
