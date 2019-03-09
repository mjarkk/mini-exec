package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Command includes data about what to exec
type Command struct {
	Cmd   string
	Args  []string
	Dir   string
	Flags map[string]string
	Kill  chan struct{}
}

// ExecWithReturnData is mostly the same as Exec, but this one also returns data
func ExecWithReturnData(c Command, newLine func(string)) error {
	cmd := exec.Command(c.Cmd, c.Args...)
	if len(c.Dir) > 0 {
		cmd.Dir = c.Dir
	}

	cmd.Env = os.Environ()
	for key, value := range c.Flags {
		cmd.Env = append(cmd.Env, key+"="+value)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		<-c.Kill
		cmd.Process.Kill()
	}()
	go func() {
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			newLine(scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			newLine(scanner.Text())
		}
	}()

	return cmd.Wait()
}

// ExecWithOutput returns the otherwhise logged data
func ExecWithOutput(c Command) (string, error) {
	data := ""
	err := ExecWithReturnData(c, func(line string) {
		prefix := "\n"
		if data == "" {
			prefix = ""
		}
		data += prefix + line
	})
	if err != nil && len(err.Error()) < 5 {
		// If the error message is so short it's probebly a error code
		return data, errors.New(data)
	}
	return data, err
}

// Exec run a command
func Exec(c Command) error {
	return ExecWithReturnData(c, func(line string) {
		fmt.Println(line)
	})
}
