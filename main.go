package main

import (
	"fmt"
  "sync"
)

type port struct {
	views int
  mu sync.Mutex
}

func (p *port) increment(w *sync.WaitGroup) {
  defer func() {
    w.Done()
    p.mu.Unlock()
  }()

  
  p.mu.Lock()
	p.views++
}

func main() {
	myPort := port{views: 0}

  var wg sync.WaitGroup
  
	for i := 0; i < 1000; i++ {
    wg.Add(1)
    go myPort.increment(&wg)
  }

  wg.Wait()


	fmt.Println(myPort.views)
}
