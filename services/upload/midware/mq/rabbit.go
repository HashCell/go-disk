package mq

import (
	"github.com/streadway/amqp"
	"go-disk/services/upload/config"
	"log"
)

var (
	channel *amqp.Channel
)

func initChannel() bool {
	if channel != nil {
		return true
	}

	conn, err := amqp.Dial(config.Conf.Mq.Rabbit.Url)
	if err != nil {
		log.Printf("failed to connect to rabbit mq server : %v", err)
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Printf("failed to get rabbit mq channel : %v", err)
		return false
	}

	return true
}

func RabbitPublish(exchange string, routingKey string, msg []byte) bool {
	if !initChannel() {
		return false
	}

	err := channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: msg,
		})

	if err != nil {
		log.Printf("publish message error : %v", err)
		return false
	}
	return true
}