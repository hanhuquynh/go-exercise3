package main

import (
	"log"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	wg sync.WaitGroup
)

func chanRoutine() {
	log.Print("hello 1")
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	wg.Wait()
	log.Print("hello 2")
}

func chanRoutineMuTex() {

	log.Print("hello 1")
	mu.Lock()
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
		mu.Unlock()
	}()
	mu.Lock()
	log.Print("hello 2")
	wg.Wait()
}

func chanRoutineChannel() {
	done := make(chan int, 1)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		done <- 1
	}()
	<-done

	log.Print("hello 2")
}
