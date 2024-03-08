package main

import (
	"fmt"
	"sync"
	"time"
)

type Baker struct{}

func (b *Baker) Bake(cakeID int, workerNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	fmt.Printf("Worker %d. Started baking cake with id: %d. Started At: %v \n", workerNumber, cakeID, startTime.String())
	time.Sleep(10 * time.Second)
	elapsed := time.Since(startTime)
	fmt.Printf("It took a total of %v time for baking cake with id: %d for worker%d\n", elapsed, cakeID, workerNumber)
	endedAt := time.Now()
	fmt.Printf("Worker %d done baking cake with id: %d. Ended At: %v \n", workerNumber, cakeID, endedAt.String())
	fmt.Println()
}

func main() {
	noOfCakes := 6

	var wg sync.WaitGroup

	startTime := time.Now()
	for i := 1; i <= noOfCakes; i++ {
		wg.Add(1)
		b := new(Baker)
		go b.Bake(i, i, &wg)
	}

	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Total time it took: %v \n", elapsed)
}
