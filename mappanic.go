package main

import (
	"os"
	"syscall"
)

var m = map[string]int{"a": 1}

//var lock = sync.RWMutex{}

func main() {
	logFile, _ := os.OpenFile("/tmp/x", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
	syscall.Dup2(int(logFile.Fd()), 1)
	syscall.Dup2(int(logFile.Fd()), 2)

	for {
		go Read()
		go Write()
	}
}

func Read() {
	for {
		read()
	}
}

func Write() {
	for {
		write()
	}
}

func read() {
	//lock.RLock()
	//defer lock.RUnlock()
	_ = m["a"]
}

func write() {
	//lock.Lock()
	//defer lock.Unlock()
	m["b"] = 2
}
