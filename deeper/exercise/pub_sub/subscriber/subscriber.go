package subscriber

import (
	"github.com/shinhagunn/go-learn/deeper/exercise/pub_sub/pubsub"
)

type PubSubProvider interface {
	RegisterSubscription(topic string) pubsub.MessageChan
}

type subscriberEngine struct {
	pb PubSubProvider
}

func NewSubscriberEngine(pb PubSubProvider) *subscriberEngine {
	return &subscriberEngine{
		pb: pb,
	}
}

func (c *subscriberEngine) Run() {
	c.subscribeTopic("update-user", updateUser)
}

func (c *subscriberEngine) subscribeTopic(topic string, handler func(message pubsub.Message)) {
	subscription := c.pb.RegisterSubscription(topic)
	go func(c pubsub.MessageChan) {
		for {
			message := <-c
			handler(message)
		}
	}(subscription)
}
