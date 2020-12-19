package main

// import "fmt"

func DrawSide(diceNum int) string {
  var diceSide string
  switch diceNum {
    case 1:
      diceSide += "\n"
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  1   ||\n"
      diceSide += "|______|/"
    case 2:
      diceSide += "\n"
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  2   ||\n"
      diceSide += "|______|/"
    case 3:
      diceSide += "\n"
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  3   ||\n"
      diceSide += "|______|/"
    case 4:
      diceSide += "\n"
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  4   ||\n"
      diceSide += "|______|/"
    case 5:
      diceSide += "\n"
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  5   ||\n"
      diceSide += "|______|/"
    case 6:
      diceSide += "\n"
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  6   ||\n"
      diceSide += "|______|/"
  }
  return diceSide
}
