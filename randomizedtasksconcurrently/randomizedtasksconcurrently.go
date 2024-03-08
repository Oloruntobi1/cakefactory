package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Baker struct {
	id         int
	tasksDone  int
	taskStream chan int
}

func (b *Baker) Bake(wg *sync.WaitGroup) {
	defer wg.Done()

	for cakeID := range b.taskStream {
		b.tasksDone++
		startTime := time.Now()
		fmt.Printf("Worker %d. Started baking cake with id: %d. Started At: %v \n", b.id, cakeID, startTime.String())
		time.Sleep(10 * time.Second)
		elapsed := time.Since(startTime)
		fmt.Printf("It took a total of %v time for baking cake with id: %d for worker%d\n", elapsed, cakeID, b.id)
		endedAt := time.Now()
		fmt.Printf("Worker %d done baking cake with id: %d. Ended At: %v \n", b.id, cakeID, endedAt.String())
	}
	fmt.Println()
	fmt.Printf("Worker %d tasks done: %d\n", b.id, b.tasksDone)
	fmt.Println()
}

func main() {
	noOfCakes := 6
	noOfWorkers := 2

	workers := make([]*Baker, 2)
	for i := 0; i < noOfWorkers; i++ {
		workers[i] = &Baker{
			id:         i + 1,
			tasksDone:  0,
			taskStream: make(chan int),
		}
	}

	var wg sync.WaitGroup

	startTime := time.Now()
	for _, worker := range workers {
		wg.Add(1)
		go worker.Bake(&wg)
	}

	// send tasks
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for cakeID := 1; cakeID <= noOfCakes; cakeID++ {
		workerIndex := rand.Intn(noOfWorkers)
		workers[workerIndex].taskStream <- cakeID
	}

	for _, worker := range workers {
		close(worker.taskStream)
	}

	wg.Wait()

	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("Total time it took: %v \n", elapsed)
}
