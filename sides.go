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

    if roll == "Y" || roll == "y" {
		  firstDice := RandomRoll()
		  secondDice := RandomRoll()
      if (firstDice == "Roll: 1" && secondDice == "Roll: 1"){
        fmt.Println("Snake eyes!")
      } else {
        fmt.Println("----------")
        fmt.Println(firstDice + "\n" + secondDice)
      }
      fmt.Println("\n----------")
      i += 1
	  } else if roll == "N" || roll == "n" {
		  fmt.Println("Goodbye!")
      i = 100
      break
	  } else {
      fmt.Println("----------\n")
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
      resultMsg = DrawSide(1)
    case 2:
      resultMsg = DrawSide(2)
    case 3:
      resultMsg = DrawSide(3)
    case 4:
      resultMsg = DrawSide(4)
    case 5:
      resultMsg = DrawSide(5)
    case 6:
      resultMsg = DrawSide(6)
	}
  return resultMsg
}