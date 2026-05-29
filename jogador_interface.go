package main

type Jogador interface {
	FazerJogada(t *Tabuleiro)
	GetNome() string
	GetSimbolo() string
	//qualquer struct que implemente essas funções
	//é um Jogador
}
