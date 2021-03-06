package service

import (
	"html/template"
	"net/http"

	"github.com/xbrett/sussex/logic"
)

//DisplayInfo is for accessing logic data
type DisplayInfo struct {
	logicDataAccess logic.Logic
	gl              logic.GroceryList
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func Movies(w http.ResponseWriter, r *http.Request) {
	var toWatch []string
	toWatch = append(toWatch, "Godfather")
	toWatch = append(toWatch, "The Other Guys")
	toWatch = append(toWatch, "Taken")
	toWatch = append(toWatch, "Get Out")

	tpl.ExecuteTemplate(w, "movies.gohtml", toWatch)
}
