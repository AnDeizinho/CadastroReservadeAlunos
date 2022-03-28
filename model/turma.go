package model
type Turma struct{
	ID int
	Serie string
	Nivel string
	Turno string
	AnoLetivo int
}

type listTurmas []Turma
var ListTurmas listTurmas

