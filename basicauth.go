package wrench

import (
	"crypto/subtle"
	"net/http"
)

// NewBasicAuth returns a one account basic auth middleware
func NewBasicAuth(account, password, realm string) *BasicAuth {
	return &BasicAuth{
		account:  account,
		password: password,
		realm:    realm,
	}
}

type BasicAuth struct {
	account  string
	password string
	realm    string
}

// BasicAuth middleware
func (me *BasicAuth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acc, pass, ok := r.BasicAuth()
		accOk := subtle.ConstantTimeCompare([]byte(acc), []byte(me.account)) == 1
		passOk := subtle.ConstantTimeCompare([]byte(pass), []byte(me.password)) == 1
		if !ok || !accOk || !passOk {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+me.realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
