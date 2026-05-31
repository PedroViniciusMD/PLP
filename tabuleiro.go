package main

import "fmt"

type Tabuleiro struct {
	casas [3][3]string //casas [i][j] onde i é a linha e j é a coluna
}

func NovoTabuleiro() *Tabuleiro {
	t := &Tabuleiro{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t.casas[i][j] = " "
		}
	}
	return t
}

func (t *Tabuleiro) Exibir() { //trabaha com o endereço do tabuleiro
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s \n", t.casas[i][0], t.casas[i][1], t.casas[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}

func (t *Tabuleiro) Marcar(linha, coluna int, simbolo string) bool {
	if t.casas[linha][coluna] != " " {
		return false
	}
	t.casas[linha][coluna] = simbolo
	return true
} //só marca se a casa estiver vazia

func (t *Tabuleiro) VerificarVencedor() string {
	//linhas e colunas
	for i := 0; i < 3; i++ {
		if t.casas[i][0] == t.casas[i][1] && t.casas[i][1] == t.casas[i][2] && t.casas[i][0] != " " {
			return t.casas[i][0]
		}
		if t.casas[0][i] == t.casas[1][i] && t.casas[1][i] == t.casas[2][i] && t.casas[0][i] != " " {
			return t.casas[0][i]
		}
	}
	//diagonais
	if t.casas[0][0] == t.casas[1][1] && t.casas[1][1] == t.casas[2][2] && t.casas[0][0] != " " {
		return t.casas[0][0]
	}
	if t.casas[0][2] == t.casas[1][1] && t.casas[1][1] == t.casas[2][0] && t.casas[0][2] != " " {
		return t.casas[0][2]
	}
	return " "
}

func (t *Tabuleiro) Empate() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.casas[i][j] == " " {
				return false
			}
		}
	}
	return true
}
