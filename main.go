package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	file1 := "Super file 1"
	var wg sync.WaitGroup
	ch := make(chan string, 4)
	file := "Super file 2"
	wg.Add(4)

	go handleApiCall(&wg, &ch)
	go handleApiCall2(&wg, &ch)
	go handleApiCall3(&wg, &ch)
	go handleApiCall0(&wg, &ch)

	close(ch)

	addedFile := file1 + file

	_ = addedFile

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

func handleApiCall0(wg *sync.WaitGroup, ch *chan string) {
	time.Sleep(time.Millisecond * 120)
	*ch <- "Fourth api call"
	wg.Done()
}
