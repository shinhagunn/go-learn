package pubsub

type Message struct {
	topic   string
	content interface{}
}

func NewMessage(topic string, content interface{}) Message {
	return Message{
		topic:   topic,
		content: content,
	}
}

func (m *Message) GetTopic() string {
	return m.topic
}

func (m *Message) GetContent() interface{} {
	return m.content
}

type MessageChan chan Message
