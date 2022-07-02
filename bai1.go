package main

import (
	"log"
	"time"
)

func chanRoutine() {
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	log.Print("hello 2")
}

func chanRoutineMuTex() {
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	log.Print("hello 2")
}
