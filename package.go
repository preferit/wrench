// Package wrench provides a web user interface for timesheet management
package wrench

import (
	"io/ioutil"

	"github.com/gregoryv/fox"
)

var DefaultLogger = fox.NewSyncLog(ioutil.Discard)
