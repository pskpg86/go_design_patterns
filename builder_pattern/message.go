//https://blog.ralch.com/articles/design-patterns/golang-builder/

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Message is the Product object in Builder Design Pattern
type Message struct {
	// Message Body
	Body []byte
	// Message Format
	Format string
}

// MessageBuilder is the inteface that every concrete implementation
// should obey
type MessageBuilder interface {
	// Set the message's recipient
	SetRecipient(recipient string)
	// Set the message's text
	SetText(text string)
	// Returns the built Message
	Message() (*Message, error)
}

// JSON Message Builder is concrete builder
type JSONMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *JSONMessageBuilder) SetRecipient(recipient string) {
	b.messageRecipient = recipient
}

func (b *JSONMessageBuilder) SetText(text string) {
	b.messageText = text
}

func (b *JSONMessageBuilder) Message() (*Message, error) {
	m := make(map[string]string)
	m["recipient"] = b.messageRecipient
	m["message"] = b.messageText

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Message{Body: data, Format: "JSON"}, nil
}

// XML Message Builder is concrete builder
type XMLMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *XMLMessageBuilder) SetRecipient(recipient string) {
	b.messageRecipient = recipient
}

func (b *XMLMessageBuilder) SetText(text string) {
	b.messageText = text
}

func (b *XMLMessageBuilder) Message() (*Message, error) {
	type XMLMessage struct {
		Recipient string `xml:"recipient"`
		Text      string `xml:"body"`
	}

	m := XMLMessage{
		Recipient: b.messageRecipient,
		Text:      b.messageText,
	}

	data, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Message{Body: data, Format: "XML"}, nil
}

// Sender is the Director in Builder Design Pattern
type Sender struct{}

// Build a concrete message via MessageBuilder
func (s *Sender) BuildMessage(builder MessageBuilder) (*Message, error) {
	builder.SetRecipient("Santa Claus")
	builder.SetText("I have tried to be good all year and hope that you and your reindeers will be able to deliver me a nice present.")
	return builder.Message()
}

func main() {
	sender := &Sender{}

	jsonMsg, err := sender.BuildMessage(&JSONMessageBuilder{})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonMsg.Body))

	xmlMsg, err := sender.BuildMessage(&XMLMessageBuilder{})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xmlMsg.Body))
}
