package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()
	// // ch := make(chan int) //un-buffered
	// ch := make(chan int, 2) // buffered
	// var wg sync.WaitGroup
	// wg.Go(func() {
	// 	fmt.Println("sending")
	// 	// This is unbuffered, so it sleeps until someone receives it.
	// 	ch <- 1
	// 	ch <- 2
	// 	// buffered 2, more than that not-consuming works like un-buffered
	// 	ch <- 3
	// 	ch <- 4 // 1 consumed, 3 want's to send but buffer allows max 2
	// 	fmt.Println("sent")
	// })
	// wg.Go(func() {
	// 	fmt.Println("receiving")
	// 	//to prove sender still waiting for someone to receive while un-buffered
	// 	// and if used buffered the sender release instantly
	// 	// as it need not to ensure someone consumed
	// 	time.Sleep(1 * time.Second)
	// 	val := <-ch
	// 	fmt.Println("Received", val)

	// })

	// wg.Wait()
	// fmt.Println("Mother goroutine Ends")
}
