package main

import "encoding/json"

func createMessageFromJson( stringJson string ) (m Message) {
	b := []byte(stringJson)
	err:=json.Unmarshal(b, &m)
	if err != nil {
		println("data is invalid")
	}
	return
}