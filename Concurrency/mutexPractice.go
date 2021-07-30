package concurrency

import (
	"fmt"
	"sync"
)

func MutexBasics() {
	raceConditionExample()
	raceConditionSolvedExample()
	raceConditionSolvedWithChannelsExample()

}

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func raceConditionExample() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

var y = 0

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}
func raceConditionSolvedExample() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementWithMutex(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}

var z = 0

func incrementWithChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}
func raceConditionSolvedWithChannelsExample() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementWithChannel(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of z", z)
}
