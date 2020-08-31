package wrench

import (
	"crypto/subtle"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	realm := "Wrench"
	r.HandleFunc("/", ServeIndex)
	r.HandleFunc("/reports", BasicAuth(ServeReports, "john", "secret", realm))
	r.HandleFunc("/reports/", BasicAuth(ServeReports, "john", "secret", realm))
	r.HandleFunc("/help", ServeHelp)
	return r
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

// BasicAuth middleware
func BasicAuth(next http.HandlerFunc, account, password, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acc, pass, ok := r.BasicAuth()
		accOk := subtle.ConstantTimeCompare([]byte(acc), []byte(account)) == 1
		passOk := subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1
		if !ok || !accOk || !passOk {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		next(w, r)
	}
}
