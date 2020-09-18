package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/fox"
	"github.com/preferit/wrench"
)

func main() {
	cli := cmdline.New(os.Args...)
	bind := cli.Option("-b, --bind").String(":8081")
	help := cli.Flag("-h, --help")

	if err := cli.Error(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if help {
		cli.WriteUsageTo(os.Stdout)
		os.Exit(0)
	}

	sl := fox.NewSyncLog(os.Stderr)
	wrench.DefaultLogger = sl

	sl.Log("listening on", bind)
	router := wrench.NewRouter()
	err := http.ListenAndServe(bind, router)
	if err != nil {
		sl.Log(err)
		os.Exit(1)
	}
}
