package main

import (
  "fmt"
  "time"
)

type DiceType struct {
  EnterMsg string
}

func (d DiceType) TimeWait() {
  fmt.Printf("%s", d.EnterMsg)
  waitTime := time.NewTimer(5 * time.Second)
  fmt.Println("\nRolling........\n")
  <- waitTime.C
}