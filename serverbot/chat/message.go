package chat

import "fmt"

type JSONMessage struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (self *JSONMessage) String() string {
	return fmt.Sprint(self.Author) + " says " + self.Body
}

func NewJSONMessage(message, author string) *JSONMessage {
	return &JSONMessage{author, message}
}

type Message struct {
	Author *string `db:"author"`
	Body   *string `db:"body"`
}

func NewMessage(author, body *string) *Message {
	if len(*author) == 0 || len(*body) == 0 {
		return &Message{
			new(string),
			new(string),
		}
	} else {
		return &Message{
			author,
			body,
		}
	}
}
