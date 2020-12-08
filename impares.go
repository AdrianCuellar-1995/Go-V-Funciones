package main

import "fmt"

func generarImpares() func() uint {
	i := uint(1)
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func main() {
	var totalImpares uint

	fmt.Println("ingresa el total de numeros impares ")
	fmt.Scanln(&totalImpares)

	if totalImpares > 0 {
		fmt.Println("---------------- ")
		sigImpar := generarImpares()
		for j := 0; j < int(totalImpares); j++ {
			fmt.Println(sigImpar())
		}
	} else {
		fmt.Println("ingresa algun valor mayor a 0")
	}
}
