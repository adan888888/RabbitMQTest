# RabbitMQ 简单演示

这是一个超级简化的RabbitMQ演示项目，专注于展示消息队列的核心概念。

## 🎯 项目结构

```
RabbitMQTest/
├── simple/                    # 简化版本
│   ├── models.go              # 数据模型
│   ├── rabbitmq.go            # RabbitMQ连接工具
│   ├── publisher_main.go      # 发布者
│   ├── subscriber_main.go     # 订阅者
│   ├── run_demo.sh           # 一键运行脚本
│   ├── 快速开始.md            # 超简单说明
│   └── README.md             # 详细说明
├── docker-compose.yml         # RabbitMQ服务
└── README.md                 # 项目说明
```

## 🚀 快速开始

### 1. 进入简化版本目录
```bash
cd simple
```

### 2. 一键运行演示
```bash
./run_demo.sh
```

### 3. 查看RabbitMQ管理界面
访问 http://localhost:15672
- 用户名: guest
- 密码: guest

## 💡 核心概念

### 发布者 (Publisher)
- 发送消息到队列
- 就像发邮件一样

### 订阅者 (Subscriber)
- 从队列接收消息
- 处理收到的消息

### 队列 (Queue)
- 消息的存储位置
- 就像邮箱一样

## 🎉 演示内容

这个演示会展示：
1. **发布者**发送3条消息到3个不同队列
2. **订阅者**同时监听3个队列并处理消息
3. 消息的完整传递过程

## 📋 预期输出

```
🚀 发布者启动
📤 发送消息到队列 'orders': 用户张三创建了订单
✅ 消息已发送到队列: orders

🎯 订阅者启动
📦 订单处理: 处理消息 ID=1, 内容=用户张三创建了订单
✅ 订单处理: 消息处理完成
```

## 🔧 技术栈

- **Go**: 编程语言
- **RabbitMQ**: 消息队列
- **Docker**: 容器化部署

## 📚 学习目标

通过这个简单演示，你将理解：
- ✅ 什么是消息队列
- ✅ 发布者-订阅者模式
- ✅ 消息的发送和接收
- ✅ 系统解耦的概念

现在就去 `simple/` 目录开始体验吧！
