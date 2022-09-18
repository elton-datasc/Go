package main

import "fmt"

func main() {
	x := 10 + 12 // expressão com operadores e operandos que retorna um resultados
	y := 5 == 5 // igualdade
	z := 5 != 5 // diferença
	fmt.Println(x, y, z) // para printar
}


// variável de escopo geral (utilizada no package) é declarada com var
// variável de escopo local (cold block) é declarada com :=
package main

import (
	"fmt"
)

var t = 10

func main() {
	x := 10 + 12            // expressão com operadores e operandos que retorna um resultados
	y := 5 == 5             // igualdade
	z := 5 != 5             // diferença
	fmt.Println(x, y, z, t) // para printar
}