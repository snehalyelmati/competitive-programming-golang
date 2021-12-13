package main

import (
	"fmt"
	"strconv"
)

func isRobotBounded(str string) bool {
	x, y := 0, 0
	angle := 90
	coordinates := ""

	for _, v := range str {
		fmt.Println(string(v))

		if string(v) == "G" {
			x, y, angle = G(x, y, angle)
		} else if string(v) == "R" {
			x, y, angle = R(x, y, angle)
		} else if string(v) == "L" {
			x, y, angle = L(x, y, angle)
		}

		coordinates = strconv.Itoa(x) + "," + strconv.Itoa(y)
		fmt.Println(coordinates, angle)
	}

	if x == 0 && y == 0 {
		return true
	}

	if angle != 90 {
		return true
	}

	return false
}

func G(x, y, angle int) (int, int, int) {
	switch angle {
	case 0:
		x++
	case 90:
		y++
	case 180:
		x--
	case 270:
		y--
	}
	return x, y, angle
}

func R(x, y, angle int) (int, int, int) {
	angle -= 90
	angle += 360
	angle %= 360
	return x, y, angle
}

func L(x, y, angle int) (int, int, int) {
	angle += 90
	angle %= 360
	return x, y, angle
}

func main() {
	commandString := "GRRRR"
	fmt.Println(isRobotBounded(commandString))
}
