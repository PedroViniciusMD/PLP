package main

type JogadorBase struct {
	nome    string
	simbolo string
}

func NovoJogadorBase(nome, simbolo string) JogadorBase {
	return JogadorBase{
		nome:    nome,
		simbolo: simbolo,
	}
}

func (j *JogadorBase) GetNome() string {
	return j.nome
}

func (j *JogadorBase) GetSimbolo() string {
	return j.simbolo
}
