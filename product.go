package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var MAX = 100
var buff = make(chan int, MAX)

// var fill = 0
// var use = 0
var chan_pro = make(chan bool, 10)
var chan_con = make(chan bool, 10)

var mutex = sync.Mutex{}

//var chan_g = make(chan bool, 1)

var wg = sync.WaitGroup{}

//var exit = make(chan bool)

func init_pro() {
	for i := 0; i < 10; i++ {
		chan_pro <- true
	}
	//chan_con <- true
	//chan_g <- true
}

func put() {
	value := rand.Intn(100)
	fmt.Println("生产： ", value)
	buff <- value
}

func get() {
	select {
	case v := <-buff:
		fmt.Println("消费: ", v)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
		//exit <- true
		wg.Done()
	}

}
func producer() {
	for {
		if ok := <-chan_pro; ok {
			mutex.Lock()
			put()
			//if o := <-chan_g; o {
			//	put()
			//}
			//chan_g <- true
			mutex.Unlock()
		}
		chan_con <- true
		//wg.Done()
	}
}

func consumer() {
	for {
		if ok := <-chan_con; ok {
			mutex.Lock()
			get()
			//if o := <-chan_g; o {
			//	get()
			//}
			//chan_g <- true
			mutex.Unlock()
		}
		chan_pro <- true
		//wg.Done()
	}
}

func main() {
	init_pro()
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go producer()
	}
	for i := 0; i < 10; i++ {
		go consumer()
	}
	wg.Wait()
	//<-exit
}
