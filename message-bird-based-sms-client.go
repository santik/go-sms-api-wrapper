package main

import (
	"github.com/messagebird/go-rest-api"
	"math/rand"
)

type MessageBirdBasedSmsClient struct {
	messagebirdClient *messagebird.Client
	UdhGenerator UdhGenerator
}

func (client MessageBirdBasedSmsClient) convertMessage(message Message) []Message {
	body := client.breakMessages(message.Message, message.max_message_length())

	if len(body) == 1 {
		return []Message{message}
	}

	return client.convertCombinedMessage(message)
}


func (client MessageBirdBasedSmsClient ) send(message Message)  {


	messages := client.convertMessage(message)

	if len(messages) == 1 {
		client.messagebirdClient.NewMessage(message.Originator, []string{message.Recipient}, message.Message, nil)
	} else {
		reference := rand.Intn(255)
		for i, messagePart := range client.convertMessage(message) {
			params := &messagebird.MessageParams{
				Type:        "binary",
				TypeDetails: messagebird.TypeDetails{"udh": client.UdhGenerator.generate(len(messages), i+1, reference)},
			}
			client.messagebirdClient.NewMessage(messagePart.Originator, []string{messagePart.Recipient}, messagePart.Message, params)
		}
	}
}
func (client MessageBirdBasedSmsClient) breakMessages(body string, maxLength int) []string {

	lenght := len(body)

	if lenght <= maxLength {
		return []string{body}
	}

	return client.splitString(body, maxLength)
}

func (client MessageBirdBasedSmsClient) splitString(s string, maxLength int) []string  {

	sRunes := []rune(s)

	stringLength := len(sRunes)

	numberOfChunks := float64(stringLength) / float64(maxLength)

	var chunks = make([]string, numberOfChunks)

	for len(sRunes) > maxLength {
		chunks = append(chunks, string(sRunes[:maxLength]))
		sRunes = sRunes[maxLength:]
	}
	if len(sRunes) > 0 {
		chunks = append(chunks, string(sRunes))
	}

	return chunks
}

func (client MessageBirdBasedSmsClient) convertCombinedMessage(message Message)[]Message {
	messageBodies := client.breakMessages(message.Message, message.max_message_length())

	var combinedMessages = make([]Message, len(messageBodies))

	for _, messagePart := range messageBodies {
		shortMessage := Message{message.Recipient, message.Originator, messagePart}
		combinedMessages = append(combinedMessages, shortMessage)
	}
	return combinedMessages
}
