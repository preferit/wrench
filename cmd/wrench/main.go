package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/preferit/wrench"
)

func main() {
	cli := cmdline.New(os.Args...)
	cli.Option("-o").String("/tmp") // deprecated, here until systemd is updated
	serve := cli.Flag("-s, --serve")
	bind := cli.Option("-b, --bind").String(":8081")
	if err := cli.Error(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if serve {
		fmt.Println("listening on", bind)
		router := wrench.NewRouter()
		log.Fatal(http.ListenAndServe(bind, router))
	}
}
