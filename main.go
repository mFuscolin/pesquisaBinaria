package main

import "fmt"

func pesquisaBinaria(lista []int, item int) int {
	baixo := 0             // 1
	alto := len(lista) - 1 // 1
	fmt.Printf("len", len(lista))
	for baixo <= alto { // 2
		meio := (baixo + alto) / 2 //3
		chute := lista[meio]
		if chute == item { // 4
			return meio
		}
		if chute > item { // 5
			alto = meio - 1
		} else { // 6
			baixo = meio + 1
		}
	}
	return 0 // 7
}

func main() {

	minha_lista := []int{1, 3, 5, 7, 9} // 8

	fmt.Println(pesquisaBinaria(minha_lista, 3)) // 9

	fmt.Println(pesquisaBinaria(minha_lista, -1)) // 10
}

// 1 baixo e alto acompanham a parte da lista que você está procurando.
// 2 Enquanto ainda não conseguiu chegar a um único eletemento...
// 3 ... verifica o elemento central
// 4 Acha o item.
// 5 O chute foi muito alto
// 6 O chute foi muito baixo
// 7 O item não existe
// 8 Vamos testá-lo
// 9 Lembre-se, as listas começam no 0. O próximo endereço índice 1.
// 10 "0" siginifica que o item não foi encontrado
