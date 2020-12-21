package main

import (
  "fmt"
  "time"
  "math/rand"
)

type DiceType struct {
  EnterMsg string
  firstDice string
  secondDice string
}

func (d DiceType) TimeWait() {
  fmt.Printf("%s", d.EnterMsg)
  waitTime := time.NewTimer(5 * time.Second)
  fmt.Println("\nRolling........\n")
  <- waitTime.C
}
func (e DiceType) ShowRollSix() {
  
  if (e.firstDice == "1" && e.secondDice == "1"){
    fmt.Println("Snake eyes!")
    } else {
      fmt.Println("----------")
      fmt.Println(e.firstDice + "\n" + e.secondDice)
    }
}
func (e DiceType) RandomSixRoll() string {
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
func (f DiceType) ShowRollTen() {
  
}