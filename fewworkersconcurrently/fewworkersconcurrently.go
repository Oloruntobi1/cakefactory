package main

import (
	"fmt"
	"sync"
	"time"
)

type Baker struct{}

func (b *Baker) Bake(cakeIDCh <-chan int, workerNumber int, wg *sync.WaitGroup) {
	defer wg.Done()

	for cakeID := range cakeIDCh {
		startTime := time.Now()
		fmt.Printf("Worker %d. Started baking cake with id: %d. Started At: %v \n", workerNumber, cakeID, startTime.String())
		time.Sleep(10 * time.Second)
		elapsed := time.Since(startTime)
		fmt.Printf("It took a total of %v time for baking cake with id: %d for worker%d\n", elapsed, cakeID, workerNumber)
		endedAt := time.Now()
		fmt.Printf("Worker %d done baking cake with id: %d. Ended At: %v \n", workerNumber, cakeID, endedAt.String())
	}
	fmt.Println()
}

func main() {
	noOfCakes := 6
	noOfWorkers := 2

	workers := make([]*Baker, 2)
	for i := 0; i < noOfWorkers; i++ {
		workers[i] = &Baker{}
	}

	var wg sync.WaitGroup

	cakeIDChan := make(chan int)

	startTime := time.Now()
	for workerID, worker := range workers {
		wg.Add(1)
		go worker.Bake(cakeIDChan, workerID+1, &wg)
	}

	// send tasks
	for cakeID := 1; cakeID <= noOfCakes; cakeID++ {
		cakeIDChan <- cakeID
	}
	close(cakeIDChan) // Close the channel after sending all cake IDs

	wg.Wait()

	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Total time it took: %v \n", elapsed)
}
