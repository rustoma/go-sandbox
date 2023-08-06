package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 4)

	wg.Add(4)

	go handleApiCall(&wg, &ch)
	go handleApiCall2(&wg, &ch)
	go handleApiCall3(&wg, &ch)
<<<<<<< HEAD
	go handleApiCall6(&wg, &ch)
=======
	go handleApiCall8(&wg, &ch)
>>>>>>> main

	close(ch)

	for val := range ch {
		log.Println(val)
	}
}

func handleApiCall(wg *sync.WaitGroup, ch *chan string) {
	time.Sleep(time.Millisecond * 80)
	*ch <- "First api call"
	wg.Done()
}

func handleApiCall2(wg *sync.WaitGroup, ch *chan string) {
	time.Sleep(time.Millisecond * 200)
	*ch <- "Second api call"
	wg.Done()
}

func handleApiCall3(wg *sync.WaitGroup, ch *chan string) {
	time.Sleep(time.Millisecond * 100)
	*ch <- "Third api call"
	wg.Done()
}

<<<<<<< HEAD
func handleApiCall6(wg *sync.WaitGroup, ch *chan string) {
=======
func handleApiCall8(wg *sync.WaitGroup, ch *chan string) {
>>>>>>> main
	time.Sleep(time.Millisecond * 120)
	*ch <- "Fourth api call"
	wg.Done()
}
