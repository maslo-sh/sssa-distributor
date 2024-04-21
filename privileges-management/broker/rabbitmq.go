package broker

import (
	"context"
	"fmt"
	"privileges-management/config"
)
import amqp "github.com/rabbitmq/amqp091-go"

func ConnectToRabbitMQServer() (*amqp.Connection, error) {
	conf := config.GetConfig()
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%d/",
			conf.RabbitMQ.AccessCredentials.Username,
			conf.RabbitMQ.AccessCredentials.Password,
			conf.RabbitMQ.Host,
			conf.RabbitMQ.Port),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	return conn, nil
}

func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}
	return ch, nil
}

func DeclareExchange(ch *amqp.Channel, exchangeName string) error {
	// Declare exchange
	err := ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}

	return nil
}

func DeclareQueue(ch *amqp.Channel, queueName string) error {
	// Declare queue
	_, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	return nil
}

// in other words: make queue interested in messages in given exchange
func CreateBind(ch *amqp.Channel, exchangeName, bindingKey, queueName string) error {
	err := ch.QueueBind(
		queueName,
		bindingKey,
		exchangeName,
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue to exchange: %v", err)
	}

	return nil
}

func PublishMessageWithRoutingKey(ch *amqp.Channel, exchangeName, routingKey, body string) error {
	ctx := context.Background()
	err := ch.PublishWithContext(ctx,
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		return fmt.Errorf("failed to bind queue to exchange: %v", err)
	}

	return nil
}

func PerformFullRabbitMQUpload(bindingKey, body string) error {
	var conn *amqp.Connection
	var channel *amqp.Channel
	var err error
	conf := config.GetConfig()
	conn, err = ConnectToRabbitMQServer()
	if err != nil {
		return err
	}
	channel, err = CreateChannel(conn)
	if err != nil {
		return err
	}
	err = DeclareExchange(channel, conf.RabbitMQ.ExchangeName)
	if err != nil {
		return err
	}
	err = DeclareQueue(channel, conf.RabbitMQ.QueueName)
	if err != nil {
		return err
	}
	err = CreateBind(channel, conf.RabbitMQ.ExchangeName, bindingKey, conf.RabbitMQ.QueueName)
	if err != nil {
		return err
	}
	err = PublishMessageWithRoutingKey(channel, conf.RabbitMQ.ExchangeName, bindingKey, body)
	if err != nil {
		return err
	}
	return nil
}
