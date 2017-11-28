# Sms Api Wrapper in Golang

This is my first(_and only_) project in Golang.
I tried to rewrite PHP sms wrapper (https://raw.githubusercontent.com/santik/sms_api_wrapper) in Go to start learning this language.

Supports MessageBird API out of the box.

**Running**

Rename config-example.json to config.json
Set correct key in config.
Set correct path is main.o::getClient()

In one terminal run

    go run *.go
    
In another terminal make a request
    
    curl localhost:8080/send -d '{"recipients":"+1234567890","originators":"originator","messagesss":"message"}' -H 'Content-Type: application/json'


Test

    go test	

**TODO**

more tests!
  
