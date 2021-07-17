package main

import (
	"fmt"
)

func main() {
	amarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := []string{"Helena", "Jonas", "José", "Juliana"}

	fmt.Println(amarelo, vermelho)

	vermelho = append(vermelho, "Luis Inácio")

	fmt.Println(vermelho)
}
