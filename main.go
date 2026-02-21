package main

import "fmt"

type Reciepient struct {
	Name string
	Email string
}

func main() {
	// unbuffered channel to hold reciepients
	recipientChan := make(chan Reciepient)
	recipients := loadReciepients("./emails.csv", recipientChan)
	fmt.Printf("Total reciepients loaded: %d\n", len(recipients))
}