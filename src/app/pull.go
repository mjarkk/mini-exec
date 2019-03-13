package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mjarkk/mini-exec/src/flags"

	gitcredentialhelper "github.com/mjarkk/go-git-http-credentials-helper"
)

// GitPull runs a git pull on the project
func GitPull() bool {
	cmd := exec.Command("git", "pull")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("[MINI-EXEC] can't get the current dir")
		return false
	}
	cmd.Dir = dir

	out, err := gitcredentialhelper.Run(cmd, func(question string) string {
		toReturn := os.Getenv("MINI_EXEC_USERNAME")
		if question == "password" {
			toReturn = os.Getenv("MINI_EXEC_PASSWORD")
		}
		if *flags.Verbose {
			fmt.Println("[MINI-EXEC] Git asked for:", question, ", responded with:", toReturn)
		}
		return toReturn
	})

	if *flags.Verbose {
		if out == nil {
			fmt.Println("[MINI-EXEC] Git pull exited with error code:", err.Error())
		} else {
			fmt.Println("[MINI-EXEC] Git pull out:\n", string(out))
		}
	}

	if err != nil {
		return false
	}

	if strings.Contains(string(out), "Already up to date") {
		return false
	}

	return true
}
