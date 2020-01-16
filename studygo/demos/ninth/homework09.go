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
	//Subscriber 订阅者，订阅者为一个管道
	Subscriber chan interface{}
	//TopicFunc 主题， 主题为一个过滤器
	TopicFunc func(v interface{}) bool
)

//Publisher 发布对象
type Publisher struct {
	lock        sync.RWMutex
	bufferSize  int
	timeout     time.Duration
	subscribers map[Subscriber]TopicFunc
}

//CreatePubliser 创建发布者
func CreatePubliser(timeout time.Duration, bufferSize int) *Publisher {
	return &Publisher{
		bufferSize:  bufferSize,
		timeout:     timeout,
		subscribers: make(map[Subscriber]TopicFunc),
	}
}

//SubscribeTopic 订阅主题
func (pub *Publisher) SubscribeTopic(topic TopicFunc) chan interface{} {
	ch := make(chan interface{}, pub.bufferSize)
	pub.lock.Lock()
	defer pub.lock.Unlock()
	pub.subscribers[ch] = topic
	return ch
}

//Subscribe 订阅所有主题
func (pub *Publisher) Subscribe() chan interface{} {
	return pub.SubscribeTopic(nil)
}

//Cancel 取消订阅
func (pub *Publisher) Cancel(sub chan interface{}) {
	pub.lock.Lock()
	defer pub.lock.Unlock()
	delete(pub.subscribers, sub)
	close(sub)
}
