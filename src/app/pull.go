package app

import (
	"os"
	"os/exec"
	"strings"

	"github.com/mjarkk/mini-exec/src/flags"
	"github.com/mjarkk/mini-exec/src/utils"

	gitcredentialhelper "github.com/mjarkk/go-git-http-credentials-helper"
)

// GitPull runs a git pull on the project
func GitPull() bool {
	cmd := exec.Command("git", "pull")
	dir, err := os.Getwd()
	if err != nil {
		utils.Println("can't get the current dir")
		return false
	}
	cmd.Dir = dir

	out, err := gitcredentialhelper.Run(cmd, func(question string) string {
		toReturn := os.Getenv("MINI_EXEC_USERNAME")
		if question == "password" {
			toReturn = os.Getenv("MINI_EXEC_PASSWORD")
		}
		if *flags.Verbose {
			utils.Println("Git asked for:", question, ", responded with:", toReturn)
		}
		return toReturn
	})

	if *flags.Verbose {
		if out == nil {
			utils.Println("Git pull exited with error code:", err.Error())
		} else {
			utils.Println("Git pull out:\n", string(out))
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
