package main

import (
	"github.com/d1once/bmi/info"
)

func main() {
	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := calculateBMI(weight, height)

	printBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
