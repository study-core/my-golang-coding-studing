package main

import (
	"fmt"
	"sync"
)

func main() {


	//taskCh := make(chan *types.Task,  20)
	//
	//sub := subcriber.NewSubcriber(taskCh)
	//
	//go sub.SendTask()
	//
	//publisher.NewPublisher(taskCh)
	//
	//
	//time.Sleep(100 * time.Second)



	errCh := make(chan error, 10)

	var wg sync.WaitGroup


	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func(wg *sync.WaitGroup, index int) {
			defer wg.Done()

			if index%2 == 0 {
				errCh <- fmt.Errorf("Gavin love kally, %d", index)
			}

		}(&wg, i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Printf("Received err: %s\n", err)
	}

	fmt.Println("Done ...")

}