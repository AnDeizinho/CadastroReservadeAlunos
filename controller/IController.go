package controller

import (
	"html/template"
	"net/http"
)
type IController struct{
	Tmpl tmpl

}
func (i * IController)index(w http.ResponseWriter, r * http.Request){
	i.Tmpl.renderizar(w, "index", "base", i)
}
func (i * IController)New(){
	if i.Tmpl == nil{
		i.Tmpl = make( map[string]*template.Template)
	}
}



