package main

import "fmt"

type JogadorX struct {
	JogadorBase //"embedding", jogadorX herda tudo que JogadorBase tem, mas não é um JogadorBase
}

func NovoJogadorX(nome string) JogadorX {
	return JogadorX{
		JogadorBase: NovoJogadorBase(nome, "X"),
	}
}

// criando o método FazerJogada que vem da interface Jogador
func (j *JogadorX) FazerJogada(t *Tabuleiro) {
	var linha, coluna int
	fmt.Printf("%s, digite uma linha e coluna (0 a 2): ", j.GetNome())
	fmt.Scan(&linha, &coluna)
	t.Marcar(linha, coluna, j.GetSimbolo())
}
