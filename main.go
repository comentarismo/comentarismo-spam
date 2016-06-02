package main

import (
	"comentarismo-spam/spamc"
	"log"
)

func main() {

	html := "<html>Hello world. I'm not a Spam, don't kill me SpamAssassin!</html>"
	client := spamc.New("127.0.0.1:783", 10)

	//the 2nd parameter is optional, you can set who (the unix user) do the call
	reply, err := client.Check(html, "saintienn")
	if err != nil {
		log.Println(err)
	}else {
		log.Println(reply.Code)
		log.Println(reply.Message)
		log.Println(reply.Vars)
	}
}