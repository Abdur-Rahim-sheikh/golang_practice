package main

import (
	"fmt"
	"sync"
	"time"
)

// import (
// 	"ecommerce/cmd"
// )

var wg sync.WaitGroup

func main() {
	// cmd.Serve()
	fnc := func(val int) {
		defer wg.Done()
		sum := 0
		for i := range val {
			for j := range i {
				sum += j
			}
		}
		fmt.Println(sum)
	}
	t1 := time.Now()
	wg.Add(1)
	go fnc(200000)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		// wg.Wait() // it will create deadlock, if above add is called
		// due to snake tail biting case
		// wg.Wait() // and if called without the above add, it will work,
		// but the result will not be returned to terminal as the waiting is done waiting :)
		if wg != nil {
			defer wg.Done()
		}

		time.Sleep(2 * time.Second)
		fmt.Println("Wow i am wake")
	}(&wg)

	wg.Add(1)
	go fnc(30000)

	wg.Wait()
	fmt.Println("All Done within: ", time.Since(t1))
}
