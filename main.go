//go:generate go run js/extract.go

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mjarkk/mini-exec/src/app"
	"github.com/mjarkk/mini-exec/src/checks"
	"github.com/mjarkk/mini-exec/src/flags"
	"github.com/mjarkk/mini-exec/src/server"

	gitcredentialhelper "github.com/mjarkk/go-git-http-credentials-helper"
)

func main() {
	server.Start()

	flag.Parse()

	gitcredentialhelper.SetupClient()

	err := checks.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *flags.OnlyCheck {
		os.Exit(0)
	}

	err = app.Init()
	if err != nil {
		fmt.Println("[MINI-EXEC] CRITICAL ERROR:", err.Error())
	}
}
