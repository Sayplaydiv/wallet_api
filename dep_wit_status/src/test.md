package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	go write(1)
	go read(21)
	go write(3)
	go read(22)
	go write(4)
	go read(23)
	go write(5)
	go read(24)
	go write(6)
	go read(25)
	go write(7)

	time.Sleep(20 * time.Second)
}

func read(i int) {
	println(i, "read start")
	m.RLock()
	var p = 0
	var pr = "read"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		println(i, pr)

	}
	m.RUnlock()
	println(i, "read end")
}

func write(i int) {
	println(i, "write start")

	m.Lock()
	var p = 0
	var pr = "write"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		println(i, pr)

	}
	m.Unlock()
	println(i, "write end")
}
