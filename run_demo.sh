#!/bin/bash

echo "=== RabbitMQ 分离式演示 ==="
echo ""

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker未运行，请先启动Docker"
    exit 1
fi

# 启动RabbitMQ
echo "🚀 启动RabbitMQ..."
docker-compose up -d

# 等待RabbitMQ启动
echo "⏳ 等待RabbitMQ启动..."
sleep 5

echo "✅ RabbitMQ已启动"
echo "📊 管理界面: http://localhost:15672 (用户名: guest, 密码: guest)"
echo ""

# 启动订阅者
echo "🎯 启动订阅者..."
cd /Users/a123123/GolandProjects/RabbitMQTest/subscriber && go run main.go &
SUBSCRIBER_PID=$!
cd /Users/a123123/GolandProjects/RabbitMQTest

# 等待订阅者启动
sleep 3

# 启动发布者
echo "📤 启动发布者..."
cd /Users/a123123/GolandProjects/RabbitMQTest/publisher && go run main.go
cd /Users/a123123/GolandProjects/RabbitMQTest

# 等待发布者完成
sleep 3

echo ""
echo "🎉 分离式演示完成！"
echo ""
echo "💡 项目结构："
echo "  - publisher/    发布者目录"
echo "  - subscriber/   订阅者目录"
echo "  - 每个目录都是独立的Go模块"
echo ""

# 停止订阅者
kill $SUBSCRIBER_PID 2>/dev/null

echo "🛑 停止服务..."
docker-compose down
