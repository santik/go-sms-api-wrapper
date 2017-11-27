package main

import (
	"github.com/messagebird/go-rest-api"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/spf13/viper"
)

func main()  {
	//curl localhost:8080 -d '{"recipient":"+1234567890","originator":"originator","message":"message"}' -H 'Content-Type: application/json'
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(writer http.ResponseWriter, request *http.Request) {

	jsonString, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Fprintf(writer, "%s", err)
	}

	m := createMessageFromJson(string(jsonString))

	client := getClient()
	client.send(m)
}

func getClient() SmsClient  {

	readConfig()

	mbClient := messagebird.New(viper.GetString("mb_key"))
	udhGenerator := UdhGenerator{}

	return MessageBirdBasedSmsClient{mbClient, udhGenerator}
}

func readConfig() {

	viper.SetConfigName("config")
	///home/alexander/go/src/go-sms-api-wrapper/
	viper.AddConfigPath("path/to/config")
	viper.ReadInConfig()
}

