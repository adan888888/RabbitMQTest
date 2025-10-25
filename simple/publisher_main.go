package main

import (
	"log"
	"time"
)

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
