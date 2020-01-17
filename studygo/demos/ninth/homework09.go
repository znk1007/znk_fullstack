package ninth

import (
	"fmt"
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

//Publish 发布信息
func (pub *Publisher) Publish(v interface{}) {
	pub.lock.Lock()
	defer pub.lock.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range pub.subscribers {
		wg.Add(1)
		go pub.serndTopic(sub, topic, v, &wg)
	}
	wg.Wait()
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

//Close 取消所有订阅
func (pub *Publisher) Close() {
	pub.lock.Lock()
	defer pub.lock.Unlock()
	for sub := range pub.subscribers {
		close(sub)
	}
}

//serndTopic 分发主题
func (pub *Publisher) serndTopic(
	sub Subscriber,
	topic TopicFunc,
	v interface{},
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	if topic != nil && !topic(v) || topic == nil {
		return
	}
	select {
	case sub <- v:
	case <-time.After(pub.timeout):
	}
}

type (
	subscriber   chan interface{}
	unSubscriber chan bool
	topicKey     interface{}
)

//MessageQueue 消息队列
type MessageQueue struct {
	lock     sync.RWMutex
	timeout  time.Duration
	messages map[topicKey]subscriber
}

//CreateMessageQueue 创建消息队列
func CreateMessageQueue(timeout time.Duration, topics ...interface{}) *MessageQueue {
	mq := &MessageQueue{
		messages: make(map[topicKey]subscriber),
		timeout:  timeout,
	}
	mq.AddTopic(topics...)
	return mq
}

//AddTopic 添加主题
func (mq *MessageQueue) AddTopic(topics ...interface{}) {
	if topics == nil || len(topics) == 0 {
		return
	}
	mq.lock.Lock()
	defer mq.lock.Unlock()
	for _, tp := range topics {
		mq.messages[tp] = make(subscriber)
	}
}

//ActiveTopic 激活主题
func (mq *MessageQueue) ActiveTopic(topic interface{}, active bool) {
	sub, ok := mq.messages[topic]
	if !ok {
		return
	}
	select {
	case data, ok := <-sub:
		fmt.Println("sub is ok: ", ok)
		fmt.Println("sub data: ", data)
	}

	// _, ok = <-sub
	// if ok && !active {
	// 	close(mq.messages[topic])
	// 	fmt.Println("inactive ok")
	// } else if !ok && active {
	// 	mq.messages[topic] = make(subscriber)
	// 	fmt.Println("active ok")
	// }
}

//Publish 发布相关主题消息
func (mq *MessageQueue) Publish(topic interface{}, v interface{}) {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if topic == nil {
		var wg sync.WaitGroup
		for _, sub := range mq.messages {
			if _, ok := <-sub; ok {
				wg.Add(1)
				go mq.handleMessages(sub, v, &wg)
			} else {
				wg.Done()
			}
		}
		wg.Wait()
		return
	}
	if sub, ok := mq.messages[topic]; ok {
		if _, ok := <-sub; ok {
			go mq.handleMessage(sub, v)
		}
	}
}

//Subscribe 订阅主题
func (mq *MessageQueue) Subscribe(topic interface{}, subFunc func(v interface{})) {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if topic == nil {
		for _, sub := range mq.messages {
			for {
				select {
				case v := <-sub:
					if subFunc != nil {
						subFunc(v)
					}
				case <-time.After(mq.timeout):
				}
			}
		}
	} else {
		if sub, ok := mq.messages[topic]; ok {
			select {
			case v := <-sub:
				if subFunc != nil {
					subFunc(v)
				}
			case <-time.After(mq.timeout):
			}
		}
	}
}

//UnSubscribe 取消订阅
func (mq *MessageQueue) UnSubscribe(topics ...interface{}) {
	mq.lock.Lock()
	mq.lock.Unlock()
	for idx, topic := range topics {
		if sub, ok := mq.messages[topic]; ok {
			close(sub)
			fmt.Println("unsub ok: ", ok)
		}
		fmt.Println("topic idx: ", idx)
	}
}

//handleMessage 处理消息群
func (mq *MessageQueue) handleMessages(sub subscriber, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case sub <- v:
	case <-time.After(mq.timeout):
	}
}

//handleMessage 处理消息
func (mq *MessageQueue) handleMessage(sub subscriber, v interface{}) {
	select {
	case sub <- v:
	case <-time.After(mq.timeout):
	}
}
