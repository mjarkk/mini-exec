package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
		if question == "password" {
			return os.Getenv("MINI_EXEC_PASSWORD")
		}
		return os.Getenv("MINI_EXEC_USERNAME")
	})

	if err != nil {
		return false
	}

	if strings.Contains(string(out), "Already up to date") {
		return false
	}

	return true
}
