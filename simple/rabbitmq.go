package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

// 简单的RabbitMQ客户端
type SimpleRabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

// 连接到RabbitMQ
func ConnectRabbitMQ() (*SimpleRabbitMQ, error) {
	// 连接到RabbitMQ服务器
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("连接失败: %v", err)
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("创建通道失败: %v", err)
	}

	return &SimpleRabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

// 关闭连接
func (r *SimpleRabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

// 发送消息
func (r *SimpleRabbitMQ) SendMessage(queueName string, message interface{}) error {
	// 声明队列
	_, err := r.channel.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 排他
		false,     // 无等待
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("声明队列失败: %v", err)
	}

	// 序列化消息
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("序列化失败: %v", err)
	}

	// 发送消息
	err = r.channel.Publish(
		"",        // 交换机
		queueName, // 路由键
		false,     // 强制
		false,     // 立即
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("发送消息失败: %v", err)
	}

	log.Printf("✅ 消息已发送到队列: %s", queueName)
	return nil
}

// 接收消息
func (r *SimpleRabbitMQ) ReceiveMessages(queueName string, handler func([]byte)) error {
	// 声明队列
	_, err := r.channel.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 排他
		false,     // 无等待
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("声明队列失败: %v", err)
	}

	// 开始消费
	msgs, err := r.channel.Consume(
		queueName, // 队列
		"",        // 消费者标签
		true,      // 自动确认
		false,     // 排他
		false,     // 无本地
		false,     // 无等待
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("开始消费失败: %v", err)
	}

	log.Printf("🎯 开始监听队列: %s", queueName)

	// 处理消息
	for msg := range msgs {
		log.Printf("📨 收到消息: %s", string(msg.Body))
		handler(msg.Body)
	}

	return nil
}
