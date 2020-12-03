package main

import (
	"sync"

	"KOHO/model"
	"KOHO/service"
)

func main() {
	input := "../input.txt"

	// channel of caller and transaction manager
	transaction := make(chan *model.Deposit, 100)

	// channel of transaction manager and producer
	response := make(chan string, 1000)

	var wg sync.WaitGroup

	wg.Add(3)

	// output generator routine
	go func() {
		service.StartProducer(response)
		wg.Done()
	}()

	// transaction manager routine
	go func() {
		service.StartManager(transaction, response)
		wg.Done()
	}()

	// transaction sender routine
	go func() {
		service.StartCaller(input, transaction)
		wg.Done()
	}()

	wg.Wait()
}
