package wrench

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/help", Help)
	return r
}

func Index(w http.ResponseWriter, r *http.Request) { indexPage().WriteTo(w) }
func Help(w http.ResponseWriter, r *http.Request)  { helpPage().WriteTo(w) }
