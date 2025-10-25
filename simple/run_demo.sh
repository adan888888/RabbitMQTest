#!/bin/bash

echo "=== RabbitMQ 简单演示 ==="
echo ""

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker未运行，请先启动Docker"
    exit 1
fi

# 启动RabbitMQ
echo "🚀 启动RabbitMQ..."
docker-compose up -d  #1.下载RabbitMQ镜像（如果本地没有）  2.创建容器并启动RabbitMQ服务   3.自动配置

# 等待RabbitMQ启动
echo "⏳ 等待RabbitMQ启动..."
sleep 5

echo "✅ RabbitMQ已启动"
echo "📊 管理界面: http://localhost:15672 (用户名: guest, 密码: guest)"
echo ""

# 启动订阅者
echo "🎯 启动订阅者..."
go run subscriber_main.go models.go rabbitmq.go &
SUBSCRIBER_PID=$!

# 等待订阅者启动
sleep 2

# 启动发布者
echo "📤 启动发布者..."
go run publisher_main.go models.go rabbitmq.go

# 等待发布者完成
sleep 3

echo ""
echo "🎉 演示完成！"
echo ""
echo "💡 你刚才看到的是："
echo "  1. 发布者发送了3条消息到不同队列"
echo "  2. 订阅者同时监听3个队列并处理消息"
echo "  3. 这就是RabbitMQ的基本用法！"
echo ""

# 停止订阅者
kill $SUBSCRIBER_PID 2>/dev/null

echo "🛑 停止服务..."
docker-compose down #停止和删除容器
