package wrench

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_routes(t *testing.T) {
	router := NewRouter()
	assert := asserter.New(t)
	exp := assert().ResponseFrom(router)
	exp.StatusCode(200, "GET", "/", nil)
	exp.StatusCode(401, "GET", "/reports/", nil)
	exp.StatusCode(401, "GET", "/reports", nil)
	exp.StatusCode(200, "GET", "/help", nil)

	cred := base64.StdEncoding.EncodeToString([]byte("john:secret"))
	auth := http.Header{"Authorization": []string{"Basic " + cred}}
	exp.StatusCode(200, "GET", "/reports", auth)
	exp.StatusCode(200, "GET", "/reports/", auth)
}
