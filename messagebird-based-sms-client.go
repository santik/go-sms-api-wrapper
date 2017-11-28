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
	body := client.breakMessageBody(message.Message, message.max_message_length())

	if len(body) == 1 {
		return []Message{message}
	}

	return client.convertToCombinedMessage(message)
}

func (client MessageBirdBasedSmsClient ) send(message Message)  {

	messages := client.convertMessage(message)

	if len(messages) == 1 {
		client.messagebirdClient.NewMessage(message.Originator, []string{message.Recipient}, message.Message, nil)
	} else {
		client.sendCombinedMessage(messages)
	}
}

func (client MessageBirdBasedSmsClient)sendCombinedMessage(messages []Message) {
	reference := rand.Intn(255)
	for i, messagePart := range messages {
		params := &messagebird.MessageParams{
			Type:        "binary",
			TypeDetails: messagebird.TypeDetails{"udh": client.UdhGenerator.generate(len(messages), i+1, reference)},
		}
		client.messagebirdClient.NewMessage(messagePart.Originator, []string{messagePart.Recipient}, messagePart.Message, params)
	}
}

func (client MessageBirdBasedSmsClient) breakMessageBody(body string, maxLength int) []string {

	lenght := len(body)

	if lenght <= maxLength {
		return []string{body}
	}

	return client.splitString(body, maxLength)
}

func (client MessageBirdBasedSmsClient) splitString(s string, maxLength int) (chunks []string)  {

	sRunes := []rune(s)

	for len(sRunes) > maxLength {
		chunks = append(chunks, string(sRunes[:maxLength]))
		sRunes = sRunes[maxLength:]
	}
	if len(sRunes) > 0 {
		chunks = append(chunks, string(sRunes))
	}
	return
}

func (client MessageBirdBasedSmsClient) convertToCombinedMessage(message Message) (combinedMessages []Message) {

	messageBodies := client.breakMessageBody(message.Message, message.max_message_length())

	for _, messagePart := range messageBodies {
		shortMessage := Message{message.Recipient, message.Originator, messagePart}
		combinedMessages = append(combinedMessages, shortMessage)
	}
	return
}
