package wrench

import (
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
}
