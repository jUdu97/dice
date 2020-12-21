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
    var roll, typeChoice string
    var firstDice, secondDice string
    
    fmt.Println("(Y/N)?")
	  fmt.Scanln(&roll)

    if roll == "Y" || roll == "y" {
      fmt.Println("\n6-sided dice........Enter 6")
      fmt.Println("\n10-sided dice........Enter 10")

      fmt.Println("What type of dice?")
      fmt.Scanln(&typeChoice)
      if (typeChoice == "6"){
        typeOne := DiceType {
          EnterMsg: "Roll a 6-sided pair of dice",
          firstDice: RandomSixRoll(),
          secondDice: RandomSixRoll(),
        }
        typeOne.TimeWait()
        typeOne.ShowRollSix()
      } else if (typeChoice == "10") {
        typeTwo := DiceType {
          EnterMsg: "Roll a 10-sided pair of dice",
        }
        firstDice = RandomTenRoll()
        secondDice = RandomTenRoll()
        typeTwo.TimeWait()
        fmt.Println("----------")
        fmt.Println(firstDice + "\n" + secondDice)
      } else {
        fmt.Println("Invalid input. Please try again.")
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

func RandomSixRoll() string {
	var resultMsg string
	var side int
  minSide := 1
  maxSide := 6
  rand.Seed(time.Now().UnixNano())
	side = rand.Intn(maxSide - minSide + 1) + minSide

	switch side {
    case 1:
      resultMsg = DrawSixSide(1)
    case 2:
      resultMsg = DrawSixSide(2)
    case 3:
      resultMsg = DrawSixSide(3)
    case 4:
      resultMsg = DrawSixSide(4)
    case 5:
      resultMsg = DrawSixSide(5)
    case 6:
      resultMsg = DrawSixSide(6)
	}
  return resultMsg
}
func RandomTenRoll() string {
	var resultTenMsg string
	var sideTen int
  minTenSide := 1
  maxTenSide := 10
  rand.Seed(time.Now().UnixNano())
	sideTen = rand.Intn(maxTenSide - minTenSide + 1) + minTenSide

	switch sideTen {
    case 1:
      resultTenMsg = DrawTenSide(1)
    case 2:
      resultTenMsg = DrawTenSide(2)
    case 3:
      resultTenMsg = DrawTenSide(3)
    case 4:
      resultTenMsg = DrawTenSide(4)
    case 5:
      resultTenMsg = DrawTenSide(5)
    case 6:
      resultTenMsg = DrawTenSide(6)
    case 7:
      resultTenMsg = DrawTenSide(7)
    case 8:
      resultTenMsg = DrawTenSide(8)
    case 9:
      resultTenMsg = DrawTenSide(9)
    case 10:
      resultTenMsg = DrawTenSide(10)
	}
  return resultTenMsg
}