package main

import (
	"net/http"

	"github.com/AnDeizinho/Crudgo/controller"
	"github.com/gorilla/mux"
)
func Rotas()*mux.Router{
	r := mux.NewRouter().StrictSlash(false)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	controller.Home(r)
	controller.Turma(r)
	return r
}
func main(){
	servidor := http.Server{Addr: ":5000", Handler: Rotas()}
	servidor.ListenAndServe()
}