package wrench

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", ServeIndex)
	r.HandleFunc("/help", ServeHelp)
	return r
}

// ServeIndex serves the root index page
func ServeIndex(w http.ResponseWriter, r *http.Request) { indexPage().WriteTo(w) }

// ServeHelp serves the help page
func ServeHelp(w http.ResponseWriter, r *http.Request) { helpPage().WriteTo(w) }
