package dashboard

import (
	"encoding/json"
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
	clients map[*websocket.Conn]string
}

func newHub() *Hub {
	return &Hub{clients: make(map[*websocket.Conn]string)}
}

func (h *Hub) add(conn *websocket.Conn, projectSlug string) {
	h.mu.Lock()
	h.clients[conn] = projectSlug
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

func (h *Hub) broadcastProject(projectSlug string, msg []byte) int {
	h.mu.Lock()
	defer h.mu.Unlock()
	count := 0
	for conn, connProject := range h.clients {
		if connProject != projectSlug {
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(h.clients, conn)
			conn.Close()
			continue
		}
		count++
	}
	return count
}

var hub = newHub()

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	projectSlug := strings.TrimPrefix(r.URL.Path, "/ws/")
	projectSlug = strings.Trim(projectSlug, "/")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	hub.add(conn, projectSlug)

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

func NavigateAll(projectSlug, path string) int {
	msg, _ := json.Marshal(map[string]string{
		"type": "navigate",
		"path": path,
	})
	return hub.broadcastProject(projectSlug, msg)
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
	for _, sub := range []string{"systems", "components", "contracts", "concepts", "flows", "decisions", "plans", "tasks", "designs", "requirements"} {
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
