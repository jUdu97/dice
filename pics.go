package main

func DrawSide(diceNum int) string {
  var diceSide string
  switch diceNum {
    case 1:
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  1   ||\n"
      diceSide += "|______|/"
    case 2:
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  2   ||\n"
      diceSide += "|______|/"
    case 3:
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  3   ||\n"
      diceSide += "|______|/"
    case 4:
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  4   ||\n"
      diceSide += "|______|/"
    case 5:
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  5   ||\n"
      diceSide += "|______|/"
    case 6:
      diceSide += " _______\n"
      diceSide += "/______/|\n"
      diceSide += "|      ||\n"
      diceSide += "|  6   ||\n"
      diceSide += "|______|/"
  }
  return diceSide
}
