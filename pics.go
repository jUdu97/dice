package main

func DrawSixSide(diceNum int) string {
  var diceSide string
  diceSide += " _______\n"
  diceSide += "/______/|\n"
  diceSide += "|      ||\n"
  diceSide += "|  "
  switch diceNum {
    case 1:
      diceSide += "1"
    case 2:
      diceSide += "2"
    case 3:
      diceSide += "3"
    case 4:
      diceSide += "4"
    case 5:
      diceSide += "5"
    case 6:
      diceSide += "6"
  }
  diceSide += "   ||\n"
  diceSide += "|______|/"
  return diceSide
}
