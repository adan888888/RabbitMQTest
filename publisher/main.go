package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// 消息结构
type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

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
		true,      // 持久化
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
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp091.Persistent, // 持久化消息
		},
	)
	if err != nil {
		return fmt.Errorf("发送消息失败: %v", err)
	}

	log.Printf("✅ 消息已发送到队列: %s", queueName)
	return nil
}

// 发布者 - 发送消息
func main() {
	log.Println("🚀 发布者启动")

	// 连接到RabbitMQ
	rabbit, err := ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer rabbit.Close()

	// 模拟发送不同类型的消息
	messages := []Message{
		{
			ID:      "1",
			Content: "用户张三创建了订单",
			Type:    "order",
		},
		{
			ID:      "2",
			Content: "订单支付成功",
			Type:    "payment",
		},
		{
			ID:      "3",
			Content: "发送邮件通知",
			Type:    "notification",
		},
	}

	// 发送消息到不同的队列
	queues := []string{"orders", "payments", "notifications"}

	for i, message := range messages {
		queueName := queues[i%len(queues)]

		log.Printf("📤 发送消息到队列 '%s': %s", queueName, message.Content)

		err := rabbit.SendMessage(queueName, message)
		if err != nil {
			log.Printf("❌ 发送失败: %v", err)
		}

		// 等待1秒再发送下一条消息
		time.Sleep(1 * time.Second)
	}

	log.Println("✅ 所有消息发送完成")
}
