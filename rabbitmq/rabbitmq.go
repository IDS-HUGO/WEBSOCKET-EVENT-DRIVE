package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	Channel   *amqp.Channel
	QueueName string
}

func NewRabbitMQConsumer() (*RabbitMQConsumer, error) {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	queueName := "NUEVA_COLA"

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Conectado a RabbitMQ, esperando mensajes...")

	return &RabbitMQConsumer{
		Channel:   ch,
		QueueName: queueName,
	}, nil
}

func (consumer *RabbitMQConsumer) ConsumeMessages() <-chan amqp.Delivery {
	log.Println("Intentando consumir mensajes de la cola:", consumer.QueueName)

	msgs, err := consumer.Channel.Consume(
		consumer.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error al consumir mensajes de RabbitMQ:", err)
	}

	log.Println("Suscrito exitosamente a la cola:", consumer.QueueName)
	return msgs
}
