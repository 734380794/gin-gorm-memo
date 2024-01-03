package mq

import (
	"gin-gorm-memo/v2/config"
	"github.com/streadway/amqp"
	"strings"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	connString := strings.Join([]string{config.RabbitMQ, "://", config.RabbitMQUser, ":", config.RabbitMQPassWord, "@", config.RabbitMQHost, ":", config.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	RabbitMq = conn
}
