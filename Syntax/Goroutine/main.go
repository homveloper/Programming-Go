package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	defer wait.Done()
}

func WaitUntil(done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	done <- true
}

func Test() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(100 * time.Millisecond)
	}
}

func ChanWrapper(fn func()) {
	done := make(chan bool)

	go func() {
		fn()
		done <- true
	}()

	<-done
}

func main() {
	wait.Add(2)
	say("① 이 루틴")
	go say("② 다른 루틴")
	wait.Wait()

	done := make(chan bool)
	go WaitUntil(done)
	<-done

	ChanWrapper(Test)
	ChanWrapper(Test)
	ChanWrapper(Test)
	ChanWrapper(Test)
}
