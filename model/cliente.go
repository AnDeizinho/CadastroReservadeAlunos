package model

type Cliente struct{
	ID int
	CPF string
	Nome string
	RendaBruta float64
	Contatos []string
	Parentesco string
	Alunos []Aluno
}

type ListClientes []Cliente
