package printinOrder

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var wg1 sync.WaitGroup

var wg2 sync.WaitGroup

func foo() {
	wg.Add(3)
	wg1.Add(1)
	wg2.Add(1)
	go first()
	go second()
	go third()
	wg.Wait()
}

func first() {
	fmt.Println("first")
	wg.Done()
	wg1.Done()
}

func second() {
	wg1.Wait()
	fmt.Println("second")
	wg.Done()
	wg2.Done()
}

func third() {
	wg2.Wait()
	fmt.Println("third")
	wg.Done()
}
