package main

import (
	"fmt"
	square "github.com/RiseLab/golang-united-school-homework-2"
)

func main() {
	println(fmt.Sprintf("%.2f", square.CalcSquare(10, square.SidesCircle)))
	println(fmt.Sprintf("%.2f", square.CalcSquare(10, square.SidesTriangle)))
	println(fmt.Sprintf("%.2f", square.CalcSquare(10, square.SidesSquare)))
	println(fmt.Sprintf("%.2f", square.CalcSquare(10, 1)))
	println(fmt.Sprintf("%.2f", square.CalcSquare(-10, square.SidesSquare)))
}
