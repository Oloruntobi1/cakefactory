package main

import (
	"fmt"
	"time"
)

type Baker struct{}

func (b *Baker) Bake(cakeID int, workerNumber int) {
	fmt.Printf("Worker %d. Started baking cake with id: %d\n", workerNumber, cakeID)
	startTime := time.Now()
	time.Sleep(10 * time.Second)
	elapsed := time.Since(startTime)
	fmt.Printf("It took a total of %v time for baking cake with id: %d for worker%d\n", elapsed, cakeID, workerNumber)
	fmt.Printf("Worker %d done baking cake with id: %d\n", workerNumber, cakeID)
	fmt.Println()
}

func main() {
	noOfCakes := 6

	startTime := time.Now()
	for i := 1; i <= noOfCakes; i++ {
		b := new(Baker)
		b.Bake(i, i)
	}

	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Total time it took: %v \n", elapsed)
}
