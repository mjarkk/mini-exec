package checks

import (
	"errors"
	"os/exec"
)

// Init will check if the needed programs are installed
func Init() error {
	_, err := exec.Command("git", "version").CombinedOutput()
	if err != nil {
		return errors.New("Git not found, this binary requires git to work")
	}
	_, err = exec.Command("git", "status").CombinedOutput()
	if err != nil {
		return errors.New("Current directory is not a git repo")
	}
	return nil
}
