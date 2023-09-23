package subscriber

import (
	"fmt"
	"time"

	"github.com/shinhagunn/go-learn/deeper/exercise/pub_sub/pubsub"
)

func updateUser(message pubsub.Message) {
	fmt.Printf("%v: Receive new message with topic: '%v' - content: %v \n", time.Now().Format("15:04:05"), message.GetTopic(), message.GetContent())
}
