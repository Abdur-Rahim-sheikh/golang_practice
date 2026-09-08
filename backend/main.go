package main

import (
	"fmt"
	"sync"
)

// import (
// 	"ecommerce/cmd"
// )

var wg sync.WaitGroup

var cnt int64
var mu sync.Mutex

func main() {
	// cmd.Serve()
	for i := 1; i <= 1000; i++ {
		wg.Go(func() {
			mu.Lock()
			a := cnt
			a = a + 1
			cnt = a
			mu.Unlock()
		})
	}

	wg.Wait()
	fmt.Println("cnt: ", cnt)
}
