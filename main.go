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
	"github.com/mjarkk/mini-exec/src/utils"

	gitcredentialhelper "github.com/mjarkk/go-git-http-credentials-helper"
)

func main() {
	flag.Parse()

	if !*flags.NoServer {
		go server.Start()
	}

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
		utils.Println("CRITICAL ERROR:", err.Error())
	}
}
