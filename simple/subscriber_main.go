package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
