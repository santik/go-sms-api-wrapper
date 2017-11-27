package main

import (
	"github.com/messagebird/go-rest-api"
	"net/http"
	"fmt"
	"io/ioutil"
)

var client SmsClient

func main()  {
	//curl localhost:8080 -d '{"recipient":"+1234567890","originator":"originator","message":"message"}' -H 'Content-Type: application/json'
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	if client == nil {
		client = getClient()
	}
}
func handler(writer http.ResponseWriter, request *http.Request) {

	jsonString, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Fprintf(writer, "%s", err)
	}

	m := createMessageFromJson(string(jsonString))

	//fmt.Println(m)

	client.send(m)
}

func getClient() SmsClient  {

	mbClient := messagebird.New("test_gshuPaZoeEG6ovbc8M79w0QyM")
	udhGenerator := UdhGenerator{}

	return MessageBirdBasedSmsClient{mbClient, udhGenerator}
}
