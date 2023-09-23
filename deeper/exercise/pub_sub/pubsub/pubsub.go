package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	locker          *sync.RWMutex
	messageQueue    MessageChan
	mapTopicMessage map[string][]MessageChan
}

func NewPubSub() *PubSub {
	return &PubSub{
		locker:          new(sync.RWMutex),
		messageQueue:    make(chan Message, 100),
		mapTopicMessage: make(map[string][]MessageChan),
	}
}

func (ps *PubSub) Broker() {
	go func() {
		for {
			message := <-ps.messageQueue

			messageChans, ok := ps.mapTopicMessage[message.GetTopic()]
			if ok {
				for _, c := range messageChans {
					c <- message
				}
			}
		}
	}()
}

func (ps *PubSub) Publish(topic string, content interface{}) {
	message := NewMessage(topic, content)

	ps.messageQueue <- message
	fmt.Printf("%v: Publish new message with topic: '%v' - content: %v\n", time.Now().Format("15:04:05"), message.GetTopic(), message.GetContent())
}

func (ps *PubSub) RegisterSubscription(topic string) MessageChan {
	newMessageChan := make(MessageChan)

	ps.locker.Lock()
	value, ok := ps.mapTopicMessage[topic]
	if ok {
		value = append(value, newMessageChan)
		ps.mapTopicMessage[topic] = value
	} else {
		ps.mapTopicMessage[topic] = []MessageChan{newMessageChan}
	}
	ps.locker.Unlock()

	return newMessageChan
}

func (ps *PubSub) Subscribe(c MessageChan) {
	go func() {
		for {
			message := <-c
			fmt.Printf("%v: Receive new message with topic: '%v' - content: %v \n", time.Now().Format("15:04:05"), message.GetTopic(), message.GetContent())
		}
	}()
}
