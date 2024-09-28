package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Dados de teste
	data := [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
	}

	for _, d := range data {
		parity := calcParity(d)

		dataWithParity := append(d, parity)

		errIndex := rand.Intn(len(dataWithParity))
		dataWithParity[errIndex] = 1 - dataWithParity[errIndex]
		detected := verifyParity(dataWithParity)

		fmt.Printf("Dados: %v\n", d)
		fmt.Printf("Bit de paridade: %d\n", parity)
		fmt.Printf("Dados com paridade: %v\n", dataWithParity)
		fmt.Printf("Erro no bit %d\n", errIndex)
		fmt.Printf("Erro detectado: %t\n\n", detected)
	}
}

func calcParity(data []int) int {
	parity := 0
	for _, bit := range data {
		parity ^= bit
	}
	return parity
}

func verifyParity(data []int) bool {
	parity := 0
	for _, bit := range data {
		parity ^= bit
	}
	return parity != 0
}

