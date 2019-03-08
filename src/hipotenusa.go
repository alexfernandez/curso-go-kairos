package main

import "fmt"
import "math"

func main() {
	fmt.Println("hipotenusa (1.5, 1.5): ", hipotenusa(1.5, 1.5))
}

func hipotenusa(cateto1 float32, cateto2 float32) float32 {
	return float32(math.Sqrt(float64(cateto1*cateto1 + cateto2*cateto2)))
}
