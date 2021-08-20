package main

import "fmt"

func main() {
	var PI float32 = 3.14
	var circleRadius int = 5
	var perimeter float32 = 2 * PI * float32(circleRadius)

	fmt.Printf("For a radius of %v, the circle perimeter is %.2f", circleRadius, perimeter)
}