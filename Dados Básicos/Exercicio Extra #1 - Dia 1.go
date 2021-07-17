package main

import (
	"fmt"
)

// 1) A função original não compila pois 150 não está dentro do range da int8
// 2) O próprio erro diz "150 overflows int8" - ou seja, ultrapassa o range de -128 até 127
// 3) Para consertar, necessitamos ou remover a atribuição int8 ou trocar outra que
// comporte o valor

func main() {

	var quilometros int
	quilometros = 150

	fmt.Println(quilometros)

}
