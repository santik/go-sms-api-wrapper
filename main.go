package main

func main()  {
	jsonString := `{"recipient":"+1234567890","originator":"originator","message":"message"}`

	m := createMessageFromJson(jsonString)

	send(m, )

	println(m.Recipient)
}
