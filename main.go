package main

import (
	"fmt"
)

type messageToSend struct {
	message		string
	sender		user
	recipient	user
}

type user struct {
	name	string
	number	int
}

func canSendMesage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	return number
}

funct main() {
	canSendMesage("Hello")
}