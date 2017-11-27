package main

type SmsSender struct {
	SmsClient SmsClient
}

func (s SmsSender)send(message Message)  {
	s.SmsClient.send(message)
}

