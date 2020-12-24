package main

func DrawSixSide(diceNum int) string {
  //Draws a 6-sided dice
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
func DrawTenSide(diceNum int) string {
  //Draws a 10-sided dice
  var diceTenSide string
  diceTenSide += " ________\n"
  diceTenSide += "/_______/|\n"
  diceTenSide += "|       ||\n"
  diceTenSide += "|  "
  switch diceNum {
    case 1:
      diceTenSide += "1"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 2:
      diceTenSide += "2"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 3:
      diceTenSide += "3"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 4:
      diceTenSide += "4"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 5:
      diceTenSide += "5"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 6:
      diceTenSide += "6"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 7:
      diceTenSide += "7"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 8:
      diceTenSide += "8"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 9:
      diceTenSide += "9"
      diceTenSide += "    ||\n"
      diceTenSide += "|_______|/"
    case 10:
      diceTenSide += "10"
      diceTenSide += "   ||\n"
      diceTenSide += "|_______|/"
  }
  return diceTenSide
}
func DrawDnDSide(diceNum int) string {
  //Draws a 20-sided dice
  var diceDnDSide string
  diceDnDSide += " ________\n"
  diceDnDSide += "/_______/|\n"
  diceDnDSide += "|       ||\n"
  diceDnDSide += "|  "
  switch diceNum {
    case 1:
      diceDnDSide += "1"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 2:
      diceDnDSide += "2"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 3:
      diceDnDSide += "3"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 4:
      diceDnDSide += "4"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 5:
      diceDnDSide += "5"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 6:
      diceDnDSide += "6"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 7:
      diceDnDSide += "7"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 8:
      diceDnDSide += "8"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 9:
      diceDnDSide += "9"
      diceDnDSide += "    ||\n"
      diceDnDSide += "|_______|/"
    case 10:
      diceDnDSide += "10"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 11:
      diceDnDSide += "11"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 12:
      diceDnDSide += "12"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 13:
      diceDnDSide += "13"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 14:
      diceDnDSide += "14"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 15:
      diceDnDSide += "15"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 16:
      diceDnDSide += "16"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 17:
      diceDnDSide += "17"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 18:
      diceDnDSide += "18"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 19:
      diceDnDSide += "19"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
    case 20:
      diceDnDSide += "20"
      diceDnDSide += "   ||\n"
      diceDnDSide += "|_______|/"
  }
  return diceDnDSide
}