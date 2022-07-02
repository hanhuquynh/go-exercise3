package main

import "fmt"

func main() {
	fmt.Println("WaitGroup")
	chanRoutine()

	fmt.Println("--------------------")

	fmt.Println("Mutex")
	chanRoutineMuTex()

	fmt.Println("--------------------")

	fmt.Println("Channel")
	chanRoutineChannel()
}
