package controllers

import(
	"net/http"
	"html/template"
)

type Home struct {
	tpl *string
}

func (h *Home) Exec(w http.ResponseWriter, req *http.Request) {
	var templ, err = template.New("qr").Parse(*(h.tpl))
	templ.Execute(w, req.FormValue("s"))

	if err != nil {
		panic(err)
	}
}

func (h *Home) SetTemplate(newTpl *string) {
	h.tpl = newTpl
}
