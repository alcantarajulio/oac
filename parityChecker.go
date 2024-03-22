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

	// Loop para testar cada conjunto de dados
	for _, d := range data {
		// Calcular o bit de paridade
		parity := calcParity(d)

		// Adicionar o bit de paridade aos dados
		dataWithParity := append(d, parity)

		// Simular um erro invertendo um bit aleatório
		errIndex := rand.Intn(len(dataWithParity))
		dataWithParity[errIndex] = 1 - dataWithParity[errIndex]

		// Verificar se o erro foi detectado
		detected := verifyParity(dataWithParity)

		// Exibir os resultados
		fmt.Printf("Dados: %v\n", d)
		fmt.Printf("Bit de paridade: %d\n", parity)
		fmt.Printf("Dados com paridade: %v\n", dataWithParity)
		fmt.Printf("Erro no bit %d\n", errIndex)
		fmt.Printf("Erro detectado: %t\n\n", detected)
	}
}

// Função para calcular o bit de paridade
func calcParity(data []int) int {
	parity := 0
	for _, bit := range data {
		parity ^= bit
	}
	return parity
}

// Função para verificar a paridade
func verifyParity(data []int) bool {
	parity := 0
	for _, bit := range data {
		parity ^= bit
	}
	return parity != 0
}

