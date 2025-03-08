// websocket_handler.go
package websocket

import (
	"log"
	"net/http"
	"websocket/rabbitmq"
)

func StartWebSocketServer() {
	http.HandleFunc("/ws", hub.HandleConnections)

	// Conectar a RabbitMQ y escuchar mensajes
	consumer, err := rabbitmq.NewRabbitMQConsumer()
	if err != nil {
		log.Fatal("Error al conectar a RabbitMQ:", err)
	}

	go func() {
		log.Println("Esperando mensajes de RabbitMQ...")
		for msg := range consumer.ConsumeMessages() {
			log.Println("Mensaje recibido de RabbitMQ:", string(msg.Body))
			hub.BroadcastMessage(string(msg.Body))
		}
		log.Println("El bucle de consumo terminó inesperadamente")
	}()

	log.Println("Servidor WebSocket escuchando en puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
