package initialize

import (
	"core/global"
	"core/tools"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var mq tools.MQ = tools.MQ{
	Exchanges: map[string]tools.Exchange{
		global.MQ_NORMAL_EXCHANGE: {
			Name:       global.MQ_NORMAL_EXCHANGE,
			Kind:       amqp.ExchangeDirect,
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		global.MQ_DLX_EXCHANGE: {
			Name:       global.MQ_DLX_EXCHANGE,
			Kind:       amqp.ExchangeDirect,
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
	},
	Queues: map[string]tools.Queue{
		global.MQ_NORMAL_QUEUE: {
			Name:       global.MQ_NORMAL_QUEUE,
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
			Binds: []tools.Bind{
				{
					Key:      "",
					Exchange: global.MQ_NORMAL_EXCHANGE,
					NoWait:   false,
					Args:     nil,
				},
			},
		},
		global.MQ_DLX_QUEUE: {
			Name:       global.MQ_DLX_QUEUE,
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args: amqp.Table{
				"x-dead-letter-exchange": global.MQ_DLX_EXCHANGE,
			},
			Binds: []tools.Bind{
				{
					Key:      "",
					Exchange: global.MQ_DLX_EXCHANGE,
					NoWait:   false,
					Args:     nil,
				},
			},
		},
	},
}

func MQ() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		log.Fatalf("MQ 链接失败：%s", err)
		return
	}

	mq.CONN = conn

	for _, exchange := range mq.Exchanges {
		if err := mq.ExchangeDeclare(exchange.Name); err != nil {
			fmt.Println("创建Exchange:", exchange.Name, "失败 err:", err)
			return
		}
		fmt.Println("创建Exchange:", exchange.Name, "成功")
	}

	for _, queue := range mq.Queues {

		if err := mq.QueueDeclare(queue.Name); err != nil {
			fmt.Println("创建Queue:", queue.Name, "失败 err:", err)
			return
		}

		fmt.Println("创建Queue:", queue.Name, "成功")
	}

	global.Mq = mq

}
