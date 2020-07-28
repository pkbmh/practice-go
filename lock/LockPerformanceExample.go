package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int64
	lock sync.Mutex
}
var wg sync.WaitGroup
func (c *Counter) increment(v int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value += v
}
func (c *Counter) getValue() int64 {
	defer c.lock.Unlock()
	c.lock.Lock()
	return c.value
}
func incr(c *Counter) {
	defer wg.Done()
	t := time.Now()
	timeout := t.Add(200*time.Millisecond)
	for t := time.Now(); t.Before(timeout); t = time.Now() {
		c.increment(1)
	}
}
func main() {
	c := &Counter{value: 0}
	for i := 0; i < 10; i = i + 1 {
		wg.Add(1)
		go incr(c)
	}
	wg.Wait()
	fmt.Println("Final value", c.value)
}
