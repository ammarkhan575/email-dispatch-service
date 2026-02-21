package main

import "sync"

type Reciepient struct {
	Name  string
	Email string
}

func main() {
	// unbuffered channel to hold reciepients
	recipientChan := make(chan Reciepient)
	var wg sync.WaitGroup
	// this will not work because the main goroutine will block on loadReciepients before emailWorker can start processing the reciepients.
	// when emailWorker start processing the reciepients, the main goroutine will be blocked on loadReciepients and will not be able to send any reciepients to the channel.
	// emailWorker(1, recipientChan)
	// loadReciepients("./emails.csv", recipientChan)

	// This will work because the emailWorker will start processing the reciepients as soon as they are loaded into the channel.
	const workerCount = 3
	go loadReciepients("./emails.csv", recipientChan)
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChan, &wg)
	}
	wg.Wait()
}
