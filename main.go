package main

import (
	"log"

	"websocket/websocket"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	websocket.StartWebSocketServer()
}
