package main

import (
	"net/http"
)

type RequestBasedSmsSender struct {
	SmsClient                  SmsClient
	requestBasedMessageCreator RequestBasedMessageCreator
}

func (sender RequestBasedSmsSender)send(writer http.ResponseWriter, request *http.Request)  {
	message := sender.requestBasedMessageCreator.create(request, writer)
	//fmt.Println(message)
	sender.SmsClient.send(message)
}
