package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

// 接收消息
func (r *SimpleRabbitMQ) ReceiveMessages(queueName string, handler func([]byte)) error {
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

// 订阅者 - 接收消息
func main() {
	log.Println("🎯 订阅者启动")

	// 连接到RabbitMQ
	rabbit, err := ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer rabbit.Close()

	// 监听订单队列
	go listenToQueue(rabbit, "orders", "📦 订单处理")

	// 监听支付队列
	go listenToQueue(rabbit, "payments", "💳 支付处理")

	// 监听通知队列
	go listenToQueue(rabbit, "notifications", "📧 通知处理")

	log.Println("✅ 所有订阅者已启动，等待消息...")
	log.Println("按 Ctrl+C 退出")

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 订阅者停止")
}

// 监听指定队列
func listenToQueue(rabbit *SimpleRabbitMQ, queueName, serviceName string) {
	err := rabbit.ReceiveMessages(queueName, func(body []byte) {
		// 解析消息
		var message Message
		if err := json.Unmarshal(body, &message); err != nil {
			log.Printf("❌ 解析消息失败: %v", err)
			return
		}

		// 处理消息
		log.Printf("%s: 处理消息 ID=%s, 内容=%s", serviceName, message.ID, message.Content)

		// 模拟处理时间
		time.Sleep(500 * time.Millisecond)

		log.Printf("✅ %s: 消息处理完成", serviceName)
	})

	if err != nil {
		log.Printf("❌ 监听队列 %s 失败: %v", queueName, err)
	}
}
