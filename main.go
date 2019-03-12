package main

import (
	"fmt"
	"os"

	"github.com/mjarkk/mini-exec/src/app"
	"github.com/mjarkk/mini-exec/src/checks"

	gitcredentialhelper "github.com/mjarkk/go-git-http-credentials-helper"
)

func main() {
	gitcredentialhelper.SetupClient()

	err := checks.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = app.Init()
	if err != nil {
		fmt.Println("[MINI-EXEC] CRITICAL ERROR:", err.Error())
	}
}
