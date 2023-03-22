package workers

import (
	"net"
	"time"

	"github.com/streadway/amqp"
)

type rabbitWorker struct {
	address     string
	connections []*rabbitConnection
	middlewares []WorkerMiddlewareFunc
}

type rabbitChannel struct {
	consumers []string
	channel   *amqp.Channel
}

type rabbitConnection struct {
	connection *amqp.Connection
	channels   []*rabbitChannel
}

func (worker *rabbitWorker) Use(newMiddlewares ...WorkerMiddlewareFunc) {
	worker.middlewares = append(worker.middlewares, newMiddlewares...)
}

func (worker *rabbitWorker) Health() error {
	conn, err := amqp.DialConfig(worker.address, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 5*time.Second)
		}})
	if err == nil {
		defer conn.Close()
	}
	return err
}

func (worker *rabbitWorker) Config(address string) error {
	worker.address = address
	return worker.Health()
}

func (worker *rabbitWorker) Close() (err error) {
	for _, connection := range worker.connections {
		if err = connection.Close(); err != nil {
			return
		}
	}
	//delete all conections after closing them
	worker.connections = []*rabbitConnection{}
	return
}

func (worker *rabbitWorker) newConnection() (conn *rabbitConnection, err error) {
	conn = new(rabbitConnection)
	conn.connection, err = amqp.Dial(worker.address)
	if err != nil {
		return
	}
	worker.connections = append(worker.connections, conn)
	return
}

func (ch rabbitChannel) Close() (err error) {
	for _, consumer := range ch.consumers {
		if err = ch.channel.Cancel(consumer, true); err != nil {
			return
		}
	}
	return ch.channel.Close()
}

func (connection rabbitConnection) Close() (err error) {
	for _, channel := range connection.channels {
		if err = channel.Close(); err != nil {
			return
		}
	}
	return connection.connection.Close()
}

func (worker *rabbitWorker) Send(queue string, messages ...[]byte) (err error) {
	var conn *rabbitConnection
	conn, err = worker.newConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	ch := new(rabbitChannel)
	ch.channel, err = conn.connection.Channel()
	if err != nil {
		return
	}
	ch.channel.QueueDeclare(queue, true, false, false, false, nil)
	for _, message := range messages {
		errPublish := ch.channel.Publish(
			"",
			queue,
			false,
			false,
			amqp.Publishing{ContentType: "application/json", Body: message},
		)
		if errPublish != nil {
			err = errPublish
		}
	}
	return err
}

func processMessage(processBody func([]byte) error, msgs <-chan amqp.Delivery, middles ...WorkerMiddleware) {
	var err error
	defer func() { recover() }()

	for d := range msgs {
		//anon func here because i want to run middleware stop no matter what
		func() {
			defer func() {
				for _, middle := range middles {
					middle.Stop(err)
				}
			}()
			for _, middle := range middles {
				middle.Start()
			}
			err = processBody(d.Body)
			if err == nil {
				//false in order to only ack this message
				d.Ack(false)
			} else {
				d.Nack(false, true)
			}
		}()
	}
}
