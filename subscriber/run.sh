#!/bin/bash

echo "=== 订阅者独立运行 ==="
echo ""

# 检查RabbitMQ是否运行
if ! docker ps | grep -q rabbitmq-server; then
    echo "🚀 启动RabbitMQ..."
    cd .. && docker-compose up -d && cd subscriber
    sleep 5
fi

echo "🎯 启动订阅者..."
go run main.go
