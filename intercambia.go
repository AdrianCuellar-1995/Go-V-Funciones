package main

import "fmt"

func intercambia(a *int64, b *int64) {
	var temp int64
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a int64
	var b int64

	fmt.Println("ingresa Valor a: ")
	fmt.Scanln(&a)
	fmt.Println("ingresa Valor b: ")
	fmt.Scanln(&b)

	intercambia(&a, &b)

	fmt.Println("ahora a es: ", a)
	fmt.Println("Ahora b es: ", b)
}
