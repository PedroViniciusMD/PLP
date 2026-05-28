package main

type JogadorX struct {
	JogadorBase //"embedding", jogadorX herda tudo que JogadorBase tem, mas não é um JogadorBase
}

func NovoJogadorX(nome string) JogadorX {
	return JogadorX{
		JogadorBase: NovoJogadorBase(nome, "X"),
	}
}
