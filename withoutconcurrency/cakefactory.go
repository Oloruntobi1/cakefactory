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
	fmt.Printf("It took a total of %v time\n", elapsed)
	fmt.Printf("Done baking cake with id: %d\n", cakeID)
}

func main() {
	noOfCakes := 6

	b := new(Baker)

	startTime := time.Now()
	for i := 1; i <= noOfCakes; i++ {
		b.Bake(i)
	}
	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Total time it took: %v \n", elapsed)
}
