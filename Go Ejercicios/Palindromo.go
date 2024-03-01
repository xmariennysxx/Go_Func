package main

import (
	"fmt"
	"strings"
)

func esPalindromo(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	var palabra string
	fmt.Print("Ingresa una palabra: ")
	fmt.Scan(&palabra)

	if esPalindromo(palabra) {
		fmt.Println("¡Es un palíndromo!")
	} else {
		fmt.Println("No es un palíndromo.")
	}
}
