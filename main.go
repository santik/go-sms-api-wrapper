package main

import (
	"github.com/messagebird/go-rest-api"
	"net/http"
	"github.com/spf13/viper"
)

func main()  {
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/queue", queueHandler)
	http.ListenAndServe(":8080", nil)
}

func sendHandler(writer http.ResponseWriter, request *http.Request) {

	smsSender := getSmsSender()
	smsSender.send(writer, request)
}

func queueHandler(writer http.ResponseWriter, request *http.Request) {

	//queueClient := getQueueClient()
	//queueClient.send(m)
}

func getSmsSender() RequestBasedSmsSender {

	readConfig()

	mbClient := messagebird.New(viper.GetString("mb_key"))
	udhGenerator := UdhGenerator{}

	client :=  MessageBirdBasedSmsClient{mbClient, udhGenerator}
	messageCreator := RequestBasedMessageCreator{}

	return RequestBasedSmsSender{client, messageCreator}
}

func readConfig() {

	viper.SetConfigName("config")
	//viper.AddConfigPath("/home/alexander/go/src/go-sms-api-wrapper/")
	viper.AddConfigPath("path/to/config")
	viper.ReadInConfig()
}

