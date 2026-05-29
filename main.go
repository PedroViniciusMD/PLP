package main

import "fmt"

type Jogo struct {
	tabuleiro    *Tabuleiro
	jogadores    [2]Jogador
	jogadorAtual int
}

func NovoJogo(j1, j2 Jogador) *Jogo {
	return &Jogo{
		tabuleiro:    NovoTabuleiro(),
		jogadores:    [2]Jogador{j1, j2},
		jogadorAtual: 0,
	}
}

func (j *Jogo) alternarJogador() {
	if j.jogadorAtual == 0 {
		j.jogadorAtual = 1
	} else {
		j.jogadorAtual = 0
	}
}

func (j *Jogo) Iniciar() {
	for {
		j.tabuleiro.Exibir()
		jogador := j.jogadores[j.jogadorAtual]
		jogador.FazerJogada(j.tabuleiro)

		vencedor := j.tabuleiro.VerificarVencedor()
		if vencedor != " " {
			j.tabuleiro.Exibir()
			fmt.Printf("%s venceu!\n", jogador.GetNome())
			break
		}

		if j.tabuleiro.Empate() {
			j.tabuleiro.Exibir()
			fmt.Println("Empate!")
			break
		}

		j.alternarJogador()
	}
}

func main() {
	x := NovoJogadorX("Pedro")
	o := NovoJogadorO("Marcus")

	jogo := NovoJogo(&x, &o)
	jogo.Iniciar()
}

// func main() {
// 	t := NovoTabuleiro()
// 	t.Exibir()
// }

// func main() {
// 	j := NovoJogadorBase("Pedro", "X")
// 	fmt.Println(j.GetNome())
// 	fmt.Println(j.GetSimbolo())
// }

// func main() {
// 	jogador_x := NovoJogadorX("Pedro")
// 	jogador_o := NovoJogadorO("Marcus")

// 	//X
// 	fmt.Println(jogador_x.GetNome())
// 	fmt.Println(jogador_x.GetSimbolo())
// 	//O
// 	fmt.Println(jogador_o.GetNome())
// 	fmt.Println(jogador_o.GetSimbolo())
// }
