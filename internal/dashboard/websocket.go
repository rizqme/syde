package dashboard

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Hub manages connected WebSocket clients and broadcasts reload events.
type Hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

func newHub() *Hub {
	return &Hub{clients: make(map[*websocket.Conn]bool)}
}

func (h *Hub) add(conn *websocket.Conn) {
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
}

func (h *Hub) remove(conn *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, conn)
	h.mu.Unlock()
	conn.Close()
}

func (h *Hub) broadcast(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for conn := range h.clients {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(h.clients, conn)
			conn.Close()
		}
	}
}

var hub = newHub()

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	hub.add(conn)

	// Keep connection alive, remove on close
	go func() {
		defer hub.remove(conn)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}

// watchSydeDir watches a .syde/ directory for changes and broadcasts reload events.
func watchSydeDir(sydeDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("syded: fsnotify error: %v", err)
		return
	}

	// Add .syde/ and subdirectories
	if err := watcher.Add(sydeDir); err != nil {
		log.Printf("syded: watch error: %v", err)
		return
	}
	// Watch entity subdirs
	for _, sub := range []string{"systems", "components", "contracts", "concepts", "flows", "decisions", "plans", "tasks", "designs", "learnings"} {
		dir := filepath.Join(sydeDir, sub)
		watcher.Add(dir) // ignore error if dir doesn't exist
	}

	// Debounce: collect events for 100ms then broadcast once
	var debounceTimer *time.Timer
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) == 0 {
				continue
			}
			// Skip index/ directory
			if strings.Contains(event.Name, "/index/") {
				continue
			}
			if debounceTimer != nil {
				debounceTimer.Stop()
			}
			debounceTimer = time.AfterFunc(100*time.Millisecond, func() {
				rel, _ := filepath.Rel(sydeDir, event.Name)
				msg := []byte(`{"type":"reload","file":"` + rel + `"}`)
				hub.broadcast(msg)
			})
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("syded: watch error: %v", err)
		}
	}
}
