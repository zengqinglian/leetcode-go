package printFooBarAlternately

import (
	"fmt"
	"sync"
)

var c1 *sync.Cond
var wg sync.WaitGroup
var isFoo bool
var lock sync.Mutex

func fooBar(n int) {
	fmt.Println("fooBar started")
	wg.Add(2)
	isFoo = true
	go foo(n)
	go bar(n)
	wg.Wait()
}

func foo(n int) {
	for i := 1; i <= n; {
		lock.Lock()
		if isFoo {
			fmt.Print("foo")
			isFoo = false
			i++
		}
		lock.Unlock()
	}
	wg.Done()
}

func bar(n int) {
	for i := 1; i <= n; {
		lock.Lock()
		if !isFoo {
			fmt.Print("bar")
			isFoo = true
			i++
		}
		lock.Unlock()
	}
	wg.Done()
}

func fooBar1(n int) {
	fmt.Println("")
	fmt.Println("fooBar1 started")
	c1 = sync.NewCond(&sync.Mutex{})
	wg.Add(2)
	isFoo = true
	go foo1(n)
	go bar1(n)
	wg.Wait()
}

func foo1(n int) {
	for i := 1; i <= n; {
		if isFoo {
			fmt.Print("foo")
			isFoo = false
			i++
		} else {
			c1.L.Lock()
			c1.Wait()
			c1.L.Unlock()
		}
		c1.Broadcast()
	}
	wg.Done()
}

func bar1(n int) {
	for i := 1; i <= n; {
		if !isFoo {
			fmt.Print("bar")
			isFoo = true
			i++
		} else {
			c1.L.Lock()
			c1.Wait()
			c1.L.Unlock()
		}
		c1.Broadcast()
	}
	wg.Done()
}
