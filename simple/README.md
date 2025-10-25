# RabbitMQ 简单演示

这是一个超级简化的RabbitMQ演示，让你快速理解消息队列的基本概念。

## 🎯 核心概念

### 发布者 (Publisher)
- 发送消息到队列
- 就像发邮件一样简单

### 订阅者 (Subscriber) 
- 从队列接收消息
- 处理收到的消息

### 队列 (Queue)
- 消息的存储位置
- 就像邮箱一样

## 📁 文件说明

```
simple/
├── models.go      # 简单的数据模型
├── rabbitmq.go    # RabbitMQ连接工具
├── publisher.go   # 发布者 - 发送消息
├── subscriber.go  # 订阅者 - 接收消息
├── run_demo.sh    # 一键运行演示
└── README.md      # 说明文档
```

## 🚀 快速体验

### 1. 启动演示
```bash
cd simple
./run_demo.sh
```

### 2. 你会看到什么
```
🚀 发布者启动
📤 发送消息到队列 'orders': 用户张三创建了订单
✅ 消息已发送到队列: orders

🎯 订阅者启动
📦 订单处理: 处理消息 ID=1, 内容=用户张三创建了订单
✅ 订单处理: 消息处理完成
```

## 💡 工作原理

```
发布者 → 队列1 → 订阅者1
发布者 → 队列2 → 订阅者2  
发布者 → 队列3 → 订阅者3
```

### 消息流程
1. **发布者**发送消息到队列
2. **订阅者**从队列接收消息
3. **订阅者**处理消息
4. 完成！

## 🔍 查看RabbitMQ管理界面

访问 http://localhost:15672
- 用户名: guest
- 密码: guest

在这里你可以看到：
- 队列列表
- 消息数量
- 连接状态

## 📝 代码解析

### 发布者代码
```go
// 发送消息到队列
err := rabbit.SendMessage(queueName, message)
```

### 订阅者代码  
```go
// 监听队列，处理消息
err := rabbit.ReceiveMessages(queueName, func(body []byte) {
    // 处理消息
})
```

## 🎉 总结

这个简单演示展示了RabbitMQ的核心功能：
- ✅ 消息发送
- ✅ 消息接收  
- ✅ 队列管理
- ✅ 解耦架构

现在你理解了RabbitMQ的基本概念！可以继续学习更复杂的应用场景。
