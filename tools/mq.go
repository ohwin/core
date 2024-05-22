package tools

import (
	"errors"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type MQ struct {
	Exchanges map[string]Exchange
	Queues    map[string]Queue
	CONN      *amqp.Connection
}

type Exchange struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

type Queue struct {
	Queue      amqp.Queue
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
	Binds      []Bind
}

type Bind struct {
	Key      string
	Exchange string
	NoWait   bool
	Args     amqp.Table
}

func (mq *MQ) Channel() (channel *amqp.Channel, err error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	channel, err = mq.CONN.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (mq *MQ) Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) (err error) {

	channel, err := mq.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer channel.Close()

	err = channel.Publish(exchange, key, mandatory, immediate, msg)
	if err != nil {
		return err
	}
	return nil
}

func (mq *MQ) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table, f func(msg amqp.Delivery)) (err error) {

	channel, err := mq.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		queue,
		consumer,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	forever := make(chan bool)

	// 处理数据
	go func() {
		for msg := range msgs {
			f(msg)
		}
	}()

	<-forever
	return nil
}

// ExchangeDeclare
// 声明交换机
func (mq *MQ) ExchangeDeclare(name string) (err error) {

	channel, err := mq.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer channel.Close()

	exchange, ok := mq.Exchanges[name]
	if !ok {
		return errors.New("exchange not found")
	}

	err = channel.ExchangeDeclare(
		exchange.Name,
		exchange.Kind,
		exchange.Durable,
		exchange.AutoDelete,
		exchange.Internal,
		exchange.NoWait,
		exchange.Args,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// QueueDeclare
// 声明队列
func (mq *MQ) QueueDeclare(name string) (err error) {

	channel, err := mq.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer channel.Close()

	queue, ok := mq.Queues[name]
	if !ok {
		return errors.New("exchange not found")
	}

	queue.Queue, err = channel.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	)
	if err != nil {
		return err
	}

	mq.Queues[name] = queue

	for _, bind := range queue.Binds {
		err = channel.QueueBind(queue.Name,
			bind.Key,
			bind.Exchange,
			bind.NoWait,
			bind.Args,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
