package controller
import(
	"html/template"
	"github.com/gorilla/mux"
)
type home struct{
	IController
}

func Home(R * mux.Router){
	var _Home home
	_Home.New()
	_Home.Tmpl["index"] = template.Must(template.ParseFiles("views/home/index.html", "views/_layaut.html"))
	R.HandleFunc("/home", _Home.index)
	R.HandleFunc("/", _Home.index)
	
	//_Home.tmpl["index"] = template.Must(template.ParseFiles("views/home/index.html", "views/_layaut.html"))
}


