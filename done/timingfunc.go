package main

import (
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "factorial")
	// ... do some things, maybe even return under some condition
	time.Sleep(10)
	return
}
