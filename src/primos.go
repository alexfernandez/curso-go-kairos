package main

import "fmt"

func main() {
	fmt.Println("suma (1000): ", suma(1000))
}

func suma(max int) int {
	suma := 0
	for i := 2; i <= max; i++ {
		if primo(i) {
			suma += i
		}
	}
	return suma
}

func primo(value int) bool {
	for i := 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
