package main

import (
	"fmt"
	"sync"
	"time"
)

// import (
// 	"ecommerce/cmd"
// )

func main() {
	// cmd.Serve()
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("sending")
		// This is unbuffered, so it sleeps until someone receives it.
		ch <- 1

		fmt.Println("sent")
	})
	wg.Go(func() {
		fmt.Println("receiving")
		time.Sleep(500 * time.Millisecond) //to prove sender still waiting for someone to receive
		val := <-ch
		fmt.Println("Received", val)

	})

	wg.Wait()
	fmt.Println("Mother goroutine Ends")
}
