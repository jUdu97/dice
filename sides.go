package main

import (
	"fmt"
	"math/rand"
  "time"
)

func DiceChoice() {
	fmt.Println("(Y/N)?")
	var roll string
	fmt.Scanln(&roll)

	if roll == "Y" {
		RandomRoll()
    fmt.Println("\t")
		RandomRoll()
	} else {
		fmt.Println("Goodbye!")
	}
}

func RandomRoll() {
	var resultMsg string
	var side int
  rand.Seed(time.Now().UnixNano())
	side = rand.Intn(6)

	switch side {
	case 1:
		resultMsg = "1"
		fmt.Print("Roll: " + resultMsg)
	case 2:
		resultMsg = "2"
		fmt.Print("Roll: " + resultMsg)
	case 3:
		resultMsg = "3"
		fmt.Print("Roll: " + resultMsg)
	case 4:
		resultMsg = "4"
		fmt.Print("Roll: " + resultMsg)
	case 5:
		resultMsg = "5"
		fmt.Print("Roll: " + resultMsg)
	case 6:
		resultMsg = "6"
		fmt.Print("Roll: " + resultMsg)
	}
}
