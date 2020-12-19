package main

import (
	"fmt"
	"math/rand"
  "time"
)

func DiceChoice() {
	i := 1
  maxRolls := 100
  for ( i < maxRolls) {
    fmt.Println("(Y/N)?")
	  var roll string
	  fmt.Scanln(&roll)

    if roll == "Y" {
		  RandomRoll()
      fmt.Println(" ")
		  RandomRoll()
      fmt.Println("\n----------")
      i += 1
	  } else if roll == "N" {
		  fmt.Println("Goodbye!")
      i = 100
      break
	  } else {
      fmt.Println("Invalid input. Try again.")
      i += 1
    }
  }	
}

func RandomRoll() {
	var resultMsg string
	var side int
  minSide := 1
  maxSide := 6
  rand.Seed(time.Now().UnixNano())
	side = rand.Intn(maxSide - minSide + 1) + minSide

	switch side {
	case 1:
		resultMsg = "1"
		fmt.Print("Roll: " + resultMsg)
    drawOne()
	case 2:
		resultMsg = "2"
		fmt.Print("Roll: " + resultMsg)
    drawTwo()
	case 3:
		resultMsg = "3"
		fmt.Print("Roll: " + resultMsg)
    drawThree()
	case 4:
		resultMsg = "4"
		fmt.Print("Roll: " + resultMsg)
    drawFour()
	case 5:
		resultMsg = "5"
		fmt.Print("Roll: " + resultMsg)
    drawFive()
	case 6:
		resultMsg = "6"
		fmt.Print("Roll: " + resultMsg)
    drawSix()
	}
}