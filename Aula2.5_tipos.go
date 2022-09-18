// Go é uma linguagem de tipagem estática (diferente de Python)
// Uma vez declarado um tipo de uma variável,este tipo so mudará se for criada uma nova varivável (Diferente de Python)

package main

import (
	"fmt"
)

var x int = 10 // variável de escopo global (declarada fora da cold box) x do tipo inteiro

func main() {

	x = 20.4 // resulta em erro pq a variável foi atualizada para um tipo diferente (float) não declarado (foi declarada com int)
	fmt.Println(x)

}