package ninth

import (
	"sync"
	"time"
)

//Product 并发对象
type Product struct {
	production chan interface{}
}

//CreateProduct 创建对象
func CreateProduct() Product {
	return Product{
		production: make(chan interface{}),
	}
}

//Produce 生产数据
func (prod Product) Produce(data interface{}) {
	go func() {
		prod.production <- data
	}()
}

// Consume 消费数据
func (prod Product) Consume(consume func(data interface{})) {
	go func() {
		for data := range prod.production {
			if consume != nil {
				consume(data)
			}
		}
	}()
}

type (
	subscriber chan interface{}         //订阅者为一个管道
	topicFunc  func(v interface{}) bool //主题为一个过滤器
)

//Publisher 发布对象
type Publisher struct {
	lock        sync.RWMutex
	bufferSize  int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

//CreatePubliser 创建发布者
func CreatePubliser(timeout time.Duration, bufferSize int) Publisher {
	return Publisher{
		bufferSize:  bufferSize,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}
