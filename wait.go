package main

import (
  "fmt"
  "time"
)

func TimeWait() {
  waitTime := time.NewTimer(5 * time.Second)
  fmt.Println("\nRolling........\n")
  <- waitTime.C
}