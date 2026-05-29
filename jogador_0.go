package main

import "fmt"

type JogadorO struct {
	JogadorBase
}

func NovoJogadorO(nome string) JogadorO {
	return JogadorO{
		JogadorBase: NovoJogadorBase(nome, "O"),
	}
}

// criando o método FazerJogada que vem da interface Jogador
func (j *JogadorO) FazerJogada(t *Tabuleiro) {
	var linha, coluna int
	fmt.Printf("%s, digite uma linha e coluna (0 a 2): ", j.GetNome())
	fmt.Scan(&linha, &coluna)
	t.Marcar(linha, coluna, j.GetSimbolo())
}
