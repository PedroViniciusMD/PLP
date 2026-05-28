package main

type JogadorO struct {
	JogadorBase
}

func NovoJogadorO(nome string) JogadorO {
	return JogadorO{
		JogadorBase: NovoJogadorBase(nome, "O"),
	}
}
