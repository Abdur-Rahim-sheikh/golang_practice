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
	ch := make(chan int) //un-buffered
	// ch := make(chan int, 2) // buffered
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("sending")
		// This is unbuffered, so it sleeps until someone receives it.
		ch <- 1

		fmt.Println("sent")
	})
	wg.Go(func() {
		fmt.Println("receiving")
		//to prove sender still waiting for someone to receive while un-buffered
		// and if used buffered the sender release instantly
		// as it need not to ensure someone consumed
		time.Sleep(1 * time.Second)
		val := <-ch
		fmt.Println("Received", val)

	})

	wg.Wait()
	fmt.Println("Mother goroutine Ends")
}
