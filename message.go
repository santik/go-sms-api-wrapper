package main

type Message struct {
	Recipient string
	Originator string
	Message string
}

func (m Message) max_message_length() int {
	return 5//160
}