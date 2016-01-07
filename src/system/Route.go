package system

import(
	"net/http"
)

type Route struct {
	Url string
	Cntr Controller
	Tpl string
}

type Controller interface {
	SetTemplate(tpl *string)
	Exec(w http.ResponseWriter, req *http.Request)
}
