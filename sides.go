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
		  firstDice := RandomRoll()
      fmt.Println(" ")
		  secondDice := RandomRoll()
      if (firstDice == "Roll: 1" && secondDice == "Roll: 1"){
        fmt.Println("Snake eyes!")
      } else {
        // fmt.Println(firstDice)
        fmt.Println(" ")
        // fmt.Println(secondDice)
      }
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

func RandomRoll() string {
	var resultMsg string
	var side int
  minSide := 1
  maxSide := 6
  rand.Seed(time.Now().UnixNano())
	side = rand.Intn(maxSide - minSide + 1) + minSide

	switch side {
    case 1:
      resultMsg = "Roll: 1" 
      drawOne()
    case 2:
      resultMsg = "Roll: 2"
      drawTwo()
    case 3:
      resultMsg = "Roll: 3"
      drawThree()
    case 4:
      resultMsg = "Roll: 4"
      drawFour()
    case 5:
      resultMsg = "Roll: 5"
      drawFive()
    case 6:
      resultMsg = "Roll: 6"
      drawSix()
	}
  return resultMsg
}