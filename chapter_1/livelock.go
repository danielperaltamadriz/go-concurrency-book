package chapter1

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func livelock() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var i atomic.Int32
	var mu chan struct{} = make(chan struct{}, 1)
	process := func(id int32) bool {
		fmt.Printf("%d :> %d acquiring\n", id, i.Load())
		i.Store(id)
		mu <- struct{}{}
		defer func() { <-mu }()
		fmt.Printf("%d :> %d waiting\n", id, i.Load())
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d :> %d comparing\n", id, i.Load())
		return i.Load() == id
	}
	wg := sync.WaitGroup{}
	fn := func(id int) {
		fmt.Printf("%d starting...\n", id)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%d :(\n", id)
				return
			default:
				ok := process(int32(id))
				if ok {
					fmt.Printf("%d :)\n", id)
					return
				}
			}
		}
	}

	for i := range 2 {
		wg.Add(1)
		go fn(i)
	}
	wg.Wait()
}
