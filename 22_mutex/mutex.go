package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc() {
	p.views++
}

func (p *post) incWithWg(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	fmt.Println(myPost.views)
	myPost.inc()
	fmt.Println(myPost.views)

	// incrementing via loop
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.incWithWg(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)

}
