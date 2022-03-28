package model
import "time"
type Aluno struct{
	ID int
	Nome string
	Novato bool
	Situacao string
	Nascimento time.Time
	Classe Turma
	Responsavel Cliente
}

type ListAlunos []Aluno
