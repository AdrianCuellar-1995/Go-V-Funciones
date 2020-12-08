package main

import "fmt"

func Nummayor(args ...int64) int64 {
	may := args[0]
	for _, valor := range args {
		if valor > may {
			may = valor
		}
	}
	return may
}

func main() {
	fmt.Println("1, 42, 2, 5, 9, 32, 76")
	fmt.Println("El número más grande es: ", Nummayor(1, 42, 2, 5, 9, 32, 76))
}
