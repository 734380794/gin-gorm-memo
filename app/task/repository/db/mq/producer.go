package mq

import (
	"fmt"
	"gin-gorm-memo/v2/consts"
	"github.com/streadway/amqp"
)

func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	mq, _ := ch.QueueDeclare(consts.RabbitMqTaskQueue, true, false, false, false, nil)
	ch.Publish("", mq.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return
	}

	fmt.Println("发送MQ成功...")
	return
}
