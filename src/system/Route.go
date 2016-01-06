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
	Exec(w http.ResponseWriter, req *http.Request)
}

