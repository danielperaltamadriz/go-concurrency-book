package chapter1

import "sync"

func racecondition() map[int]bool {
	mu := make(chan struct{}, 1)
	out := make(map[int]bool)
	var wg sync.WaitGroup

	fn := func(v bool) {
		defer wg.Done()
		mu <- struct{}{}
		out[1] = v
		<-mu
	}
	for i := range 5 {
		wg.Add(1)
		go fn(i%2 == 0)
	}

	wg.Wait()
	return out
}
