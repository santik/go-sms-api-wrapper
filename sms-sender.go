package main

func send(message Message, client SmsClient)  {
	client.send(message)
}
