package main

type SmsClient interface {
	send(message Message) float64
}