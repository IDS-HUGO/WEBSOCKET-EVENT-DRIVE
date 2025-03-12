package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan string
}

var hub = Hub{
	Clients:   make(map[*websocket.Conn]bool),
	Broadcast: make(chan string),
}

func (h *Hub) HandleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error de conexión:", err)
		return
	}
	defer conn.Close()

	// Log cuando un cliente se conecta
	log.Println("Nuevo cliente conectado:", conn.RemoteAddr())

	// Agregar el cliente al hub
	h.Clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error leyendo mensaje:", err)
			delete(h.Clients, conn)
			break
		}
		log.Printf("Mensaje recibido de %s: %s", conn.RemoteAddr(), msg)

	}
}

func (h *Hub) BroadcastMessage(message string) {
	log.Printf("Broadcasting message: %s", message)

	for client := range h.Clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error enviando mensaje:", err)
			client.Close()
			delete(h.Clients, client)
		}
	}
}
