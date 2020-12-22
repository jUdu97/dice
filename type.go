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
func (d DiceType) ShowRollDnD() {
  d.firstDice = d.RandomDnDRoll()
  fmt.Println("----------")
  fmt.Println(d.firstDice)
}
func (d DiceType) RandomDnDRoll() string {
	var resultDnDMsg string
	var sideDnD int
  minTenSide := 1
  maxTenSide := 20
  rand.Seed(time.Now().UnixNano())
	sideDnD = rand.Intn(maxTenSide - minTenSide + 1) + minTenSide

	switch sideDnD {
    case 1:
      resultDnDMsg = DrawDnDSide(1)
    case 2:
      resultDnDMsg = DrawDnDSide(2)
    case 3:
      resultDnDMsg = DrawDnDSide(3)
    case 4:
      resultDnDMsg = DrawDnDSide(4)
    case 5:
      resultDnDMsg = DrawDnDSide(5)
    case 6:
      resultDnDMsg = DrawDnDSide(6)
    case 7:
      resultDnDMsg = DrawDnDSide(7)
    case 8:
      resultDnDMsg = DrawDnDSide(8)
    case 9:
      resultDnDMsg = DrawDnDSide(9)
    case 10:
      resultDnDMsg = DrawDnDSide(10)
    case 11:
      resultDnDMsg = DrawDnDSide(11)
    case 12:
      resultDnDMsg = DrawDnDSide(12)
    case 13:
      resultDnDMsg = DrawDnDSide(13)
    case 14:
      resultDnDMsg = DrawDnDSide(14)
    case 15:
      resultDnDMsg = DrawDnDSide(15)
    case 16:
      resultDnDMsg = DrawDnDSide(16)
    case 17:
      resultDnDMsg = DrawDnDSide(17)
    case 18:
      resultDnDMsg = DrawDnDSide(18)
    case 19:
      resultDnDMsg = DrawDnDSide(19)
    case 20:
      resultDnDMsg = DrawDnDSide(20)
	}
  return resultDnDMsg
}