package main

import "github.com/messagebird/go-rest-api"

func main()  {
	jsonString := `{"recipient":"+1234567890","originator":"originator","message":"message message me"}`

	m := createMessageFromJson(jsonString)

	client := getClient()

	client.send(m)

}

func getClient() SmsClient  {

	mbClient := messagebird.New("test_gshuPaZoeEG6ovbc8M79w0QyM")
	udhGenerator := UdhGenerator{}

	return MessageBirdBasedSmsClient{mbClient, udhGenerator}
}
