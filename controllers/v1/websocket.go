package v1

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct{}

type Message struct {
	RoomID string `json:"roomID"`
	Text   string `json:"text"`
}

type Room struct {
	ID        string
	Clients   map[*websocket.Conn]bool
	Broadcast chan Message
	Lock      sync.Mutex
}

var rooms = make(map[string]*Room)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Check if the origin is allowed
		allowedOrigins := []string{"http://localhost:3000"} // Add your origin(s) here
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				return true
			}
		}
		return false
	},
}

func (ws *WebSocketController) HandleWebSocket(c *gin.Context) {
	origin := c.GetHeader("Origin")
	log.Println("WebSocket Origin:", origin)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	roomID := c.Query("roomID")
	if roomID == "" {
		log.Println("Missing roomID in query parameters")
		conn.Close()
		return
	}
	room, exists := rooms[roomID]
	room.Lock.Lock()
	if exists {
		if len(room.Clients) >= 4 {
			log.Println("Room is full")
			conn.Close()
			return
		}
	} else {
		// Create a new room instance
		room = createRoom(roomID)
		rooms[roomID] = room
	}
	room.Clients[conn] = true
	room.Lock.Unlock()

	defer func() {
		room.Lock.Lock()
		delete(room.Clients, conn)
		room.Lock.Unlock()
		conn.Close()
	}()

	go func() {
		for {
			var message Message
			err := conn.ReadJSON(&message)
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}

			message.RoomID = roomID
			room.Broadcast <- message
		}
	}()
}

func createRoom(roomID string) *Room {
	room := &Room{
		ID:        roomID,
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan Message),
	}
	log.Println("Created room:", roomID)
	go func() {
		for {
			message := <-room.Broadcast

			room.Lock.Lock()
			for client := range room.Clients {
				err := client.WriteJSON(message)
				if err != nil {
					log.Println("WebSocket write error:", err)
					client.Close()
					delete(room.Clients, client)
				}
			}
			room.Lock.Unlock()
		}
	}()

	return room
}
