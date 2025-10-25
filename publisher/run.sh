#!/bin/bash

echo "=== 发布者独立运行 ==="
echo ""

# 检查RabbitMQ是否运行
if ! docker ps | grep -q rabbitmq-server; then
    echo "🚀 启动RabbitMQ..."
    cd .. && docker-compose up -d && cd publisher
    sleep 5
fi

echo "📤 启动发布者..."
go run main.go
