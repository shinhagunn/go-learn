package main

import (
	"time"

	"github.com/shinhagunn/go-learn/deeper/exercise/pub_sub/pubsub"
	"github.com/shinhagunn/go-learn/deeper/exercise/pub_sub/subscriber"
)

const topic = "update-user"

func main() {
	pb := pubsub.NewPubSub()
	pb.Broker()

	subscribeEngine := subscriber.NewSubscriberEngine(pb)
	subscribeEngine.Run()

	pb.Publish(topic, "Hello cac ban")

	time.Sleep(5 * time.Minute)
}
