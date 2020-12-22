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
func (d DiceType) ShowRollSix() {
  d.firstDice = d.RandomSixRoll()
  d.secondDice = d.RandomSixRoll()
  if (d.firstDice == "1" && d.secondDice == "1"){
    fmt.Println("Snake eyes!")
  } else {
      fmt.Println("----------")
      fmt.Println(d.firstDice + "\n" + d.secondDice)
  }
}
func (d DiceType) RandomSixRoll() string {
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
func (d DiceType) ShowRollTen() {
  d.firstDice = d.RandomTenRoll()
  d.secondDice = d.RandomTenRoll()
  fmt.Println("----------")
  fmt.Println(d.firstDice + "\n" + d.secondDice)
}
func (d DiceType) RandomTenRoll() string {
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