package main

// 简单的消息结构
type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"` // "order", "payment", "notification"
}

// 订单信息
type Order struct {
	ID       string  `json:"id"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
