package chapter1

import "sync"

func datarace() {

	noSafeMap := make(map[int]struct{})

	wg := sync.WaitGroup{}
	fn := func(v int) {
		defer wg.Done()
		noSafeMap[v] = struct{}{}
	}

	for i := range 2 {
		wg.Add(1)
		go fn(i)
	}

	wg.Wait()
}
