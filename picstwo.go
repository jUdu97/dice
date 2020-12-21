package main

func DrawTenSide(diceNum int) string {
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