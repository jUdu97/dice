package main

import (
	"fmt"
)

func DiceChoice() {
	i := 1
  maxRolls := 100
  for ( i < maxRolls) {
    var roll, typeChoice string
    
    fmt.Println("(Y/N)?")
	  fmt.Scanln(&roll)

    if (roll == "Y" || roll == "y") {
      fmt.Println("\n6-sided dice........Enter 6")
      fmt.Println("\n10-sided dice........Enter 10")
      fmt.Println("\n20-sided dice........Enter 20")

      fmt.Println("What type of dice?")
      fmt.Scanln(&typeChoice)
      if (typeChoice == "6"){
        typeOne := DiceType {}
        typeOne.ShowRollSix()
      } else if (typeChoice == "10") {
        typeTwo := DiceType {}
        typeTwo.ShowRollTen()
      } else if (typeChoice == "20") {
        typeThree := DiceType {}
        typeThree.ShowRollDnD()
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
      fmt.Println("----------")
      fmt.Println("Invalid input. Try again.")
      i += 1
    }
  }	
}