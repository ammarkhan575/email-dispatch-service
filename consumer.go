package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Reciepient, wg *sync.WaitGroup) {
	defer wg.Done()
	for reciepient := range ch {
		fmt.Printf("Worker %d processing reciepient: %s <%s>\n", id, reciepient.Name, reciepient.Email)
	}
}