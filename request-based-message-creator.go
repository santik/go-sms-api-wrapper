package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type RequestBasedMessageCreator struct {
}

func (creator RequestBasedMessageCreator) create(request *http.Request, writer http.ResponseWriter) Message {
	jsonString, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	return creator.createMessageFromJson(string(jsonString))
}

func (RequestBasedMessageCreator) createMessageFromJson( stringJson string ) (m Message) {
	b := []byte(stringJson)
	err:=json.Unmarshal(b, &m)
	if err != nil || m.Message == "" || m.Recipient == "" || m.Originator == "" {
		panic("data is invalid")
	}
	return
}