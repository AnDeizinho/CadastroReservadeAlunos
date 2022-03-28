package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"github.com/AnDeizinho/Crudgo/model"
	"github.com/gorilla/mux"
)
type turma struct{
	IController
}
func Turma(R * mux.Router){
	var t turma
	t.New()
	t.Tmpl["index"] = template.Must(template.ParseFiles("views/turma/index.html", "views/_layaut.html"))
	t.Tmpl["turma/cad"]=template.Must(template.ParseFiles("views/turma/cadastrar.html", "views/_layaut.html"))
	t.Tmpl["turma/list"]=template.Must(template.ParseFiles("views/turma/listar.html", "views/_layaut.html"))
	R.HandleFunc("/turmas", t.index)
	R.HandleFunc("/turmas-cadastrar", t.Cadastro)
	R.HandleFunc("/turmas-cad", t.Cad)
	R.HandleFunc("/turmas-list", t.Listar)
}
func (t * turma)Cadastro(w http.ResponseWriter, r * http.Request){
	t.Tmpl.renderizar(w, "turma/cad", "base", t)
}
func (t * turma)Cad(w http.ResponseWriter, r * http.Request){
	r.ParseForm()
	novo :=model.Turma{}
	novo.Nivel = r.PostFormValue("nivel")
	novo.Serie = r.PostFormValue("serie")
	novo.Turno = r.PostFormValue("turno")
	novo.AnoLetivo , _= strconv.Atoi(r.PostFormValue("ano"))
	model.ListTurmas = append(model.ListTurmas, novo)
	http.Redirect(w, r ,"/turmas-cadastrar", 302)
}
func (t * turma)Listar(w http.ResponseWriter, r * http.Request){
	t.Tmpl.renderizar(w, "turma/list","base", model.ListTurmas)
}
