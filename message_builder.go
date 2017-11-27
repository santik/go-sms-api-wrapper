package main

import "encoding/json"

func createMessageFromJson( stringJson string ) (m Message) {
	b := []byte(stringJson)
	err:=json.Unmarshal(b, &m)
	if err != nil || m.Message == "" || m.Recipient == "" || m.Originator == "" {
		panic("data is invalid")
	}
	return
}