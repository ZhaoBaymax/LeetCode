package main

import (
	"fmt"
	"sync"
)

// 1. 使用chanel控制两个协程交替执行，并由chanel控制主进程的结束
var m1_chan1 = make(chan bool, 1)
var m1_chan2 = make(chan bool)
var m1_index = make(chan bool, 1)

func m1_1() {
	for i := 0; i < 10; i += 2 {
		if ok := <-m1_chan1; ok {
			fmt.Println(i)
			m1_chan2 <- true
		}

	}
}

func m1_2() {
	for i := 1; i < 11; i += 2 {
		if ok := <-m1_chan2; ok {
			fmt.Println(i)
			m1_chan1 <- true
		}
	}
	m1_index <- true
}

func m1() {
	m1_chan1 <- true
	go m1_1()
	go m1_2()
	<-m1_index
}

// 2. 通过channl控制两个循环交替打印, 使用sync.WaitGroup方法来控制等待主进程
var m2_chan1 = make(chan bool, 1)
var m2_chan2 = make(chan bool)
var m2_wg sync.WaitGroup

func m2_1() {
	defer m2_wg.Done()
	for i := 0; i < 10; i += 2 {
		if ok := <-m2_chan1; ok {
			fmt.Println(i)
		}
		m2_chan2 <- true
	}
}
func m2_2() {
	defer m2_wg.Done()
	for i := 1; i < 10; i += 2 {
		if ok := <-m2_chan2; ok {
			fmt.Println(i)
		}
		m2_chan1 <- true
	}
}
func m2() {
	m2_wg.Add(2)
	m2_chan1 <- true
	go m2_1()
	go m2_2()
	m2_wg.Wait()
}

var m3_wg sync.WaitGroup

func m3() {
	var m3_chan1 = make(chan bool, 1)
	//var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		m3_wg.Add(1)
		//mutex.Lock()
		m3_chan1 <- true
		go func(i int) {
			fmt.Println(i)
			//mutex.Unlock()
			<-m3_chan1
			m3_wg.Done()
		}(i)
	}
}

func main() {
	m3()
	m3_wg.Wait()
}
