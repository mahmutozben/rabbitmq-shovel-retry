package main

import (
	"fmt"
	"gitlab.hepsiburada.com/campaign/rabbitmq-shovel-retry/rabbitClient"
	"gitlab.hepsiburada.com/campaign/rabbitmq-shovel-retry/slackClient"
	"os"
)

func main() {
	fmt.Println("hello rabbit")
	//os.Setenv("ENV","qa")
	env := os.Getenv("ENV")
	fmt.Println(fmt.Sprintf("ENV :%s",env))
	config, err := GetConfig(env)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rClient := rabbitClient.NewRabbitClient(config.RabbitMq)
	queueList := GetQueueList()
	for _, qname := range queueList {
		errorQueue := fmt.Sprintf("%s_error",qname)
		queueInfo := rClient.GetQueueInfo(errorQueue)
		if queueInfo != nil && queueInfo.Messages > 0 {
			message := fmt.Sprintf("Server: '%s', RabbitMQ: '%s' işlenemeyen mesaj sayısı : '%d', Consumer sayısı : '%d' @here Errorların taşıma işlemi yapıldı :fireball:",
				config.RabbitMq.Server, errorQueue, queueInfo.Messages, queueInfo.Consumers)
			result, err := rClient.MoveErrorToQueue(qname)
			if err != nil {
				fmt.Errorf(err.Error())
			}
			if result {
				fmt.Println(message)
				slackClient.SendNotification(message, config.ChannelName)
			}
		}
	}
	fmt.Println("Job completed.")
}


