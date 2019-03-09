package app

import (
	"strings"

	"github.com/mjarkk/mini-exec/src/commands"
)

// GitPull runs a git pull on the project
func GitPull() bool {
	out, err := commands.ExecWithOutput(commands.Command{
		Cmd:  "git",
		Args: []string{"pull"},
		Flags: map[string]string{
			"LANG":   "en_US.UTF-8",
			"LC_ALL": "en_US.UTF-8",
		},
	})
	if err != nil {
		return false
	}
	if strings.Contains(out, "Already up to date") {
		return false
	}
	return true
}
