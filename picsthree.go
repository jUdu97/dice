package main

func DrawTwentySide(diceNum int) string {
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