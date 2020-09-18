package wrench

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *Router {
	r := Router{mux.NewRouter()}
	auth := NewBasicAuth("john", "secret", "Wrench")

	r.HandleFunc("/", ServeIndex)
	r.WrapFunc("/reports", ServeReports, auth)
	r.WrapFunc("/reports/", ServeReports, auth)
	r.HandleFunc("/help", ServeHelp)
	return &r
}

type Router struct {
	*mux.Router
}

// WrapFunc registers the path with middlewares prepended if any are given
func (me *Router) WrapFunc(path string, h http.HandlerFunc, mw ...middleware) {
	var handler http.Handler = http.HandlerFunc(h)
	if len(mw) > 0 {
		for l := len(mw) - 1; l >= 0; l-- {
			handler = mw[l].Middleware(handler)
		}
	}
	me.Router.Handle(path, handler)
}

type middleware interface {
	Middleware(http.Handler) http.Handler
}

// ServeIndex serves the root index page
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	view := NewIndexView()
	view.Render().WriteTo(w)
}

// ServeReports serves the reports front page
func ServeReports(w http.ResponseWriter, r *http.Request) {
	acc, _, _ := r.BasicAuth()
	view := NewReportsView(acc)
	view.Reports = append(view.Reports, Report{Text: "raw report here..."})
	view.Render().WriteTo(w)
}

// ServeHelp serves the help page
func ServeHelp(w http.ResponseWriter, r *http.Request) {
	view := NewHelpView()
	view.Render().WriteTo(w)
}
