package main

import (
	"fmt"
	"time"
)

type Baker struct{}

func (b *Baker) Bake(cakeID int) {
	fmt.Printf("Started baking cake with id: %d\n", cakeID)
	startTime := time.Now()
	time.Sleep(10 * time.Second)
	elapsed := time.Since(startTime)
	fmt.Printf("It took a total of %v time for baking cake with id: %d\n", elapsed, cakeID)
	fmt.Printf("Done baking cake with id: %d\n", cakeID)
	fmt.Println()
	// done <- struct{}{}
}

func main() {
	noOfCakes := 6

	b1 := new(Baker)

	startTime := time.Now()
	for i := 1; i <= noOfCakes-3; i++ {
		b1.Bake(i)
	}

	b2 := new(Baker)
	for i := 4; i <= noOfCakes; i++ {
		b2.Bake(i)
	}
	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Total time it took: %v \n", elapsed)
}
