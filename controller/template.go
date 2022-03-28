package controller
import(
	"net/http"
	"html/template"
)
type tmpl map[string]*template.Template

func (t tmpl)renderizar(w http.ResponseWriter, chave string, template string, dado interface{}){
	temp, ok := t[chave]
	if !ok{
		http.Error(w, "deu ruim no template", http.StatusInternalServerError)
	}
	er := temp.ExecuteTemplate(w, template, dado)
	if er !=nil {
		http.Error(w, "erro ao executar template", http.StatusInternalServerError)
	}
}
