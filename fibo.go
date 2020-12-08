package main

import "fmt"

func Fibo(n uint64) uint64 {
	if n == 0 || n == 1 {
		return n
	} else {
		return Fibo(n-1) + Fibo(n-2)
	}
}

func main() {
	var pos uint64
	fmt.Println("ingresa el numero limite a visualizar ")
	fmt.Scanln(&pos)
	if pos == 0 {
		fmt.Println("Fibonacci(", pos, ") = ", Fibo(uint64(pos)))
	} else {
		for i := 0; i < int(pos)+1; i++ {
			fmt.Println(i, ")  ", Fibo(uint64(i)))
		}
	}
}
