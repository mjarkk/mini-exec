// Package files handles the parsing of the .miniex files
package files

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/mjarkk/mini-exec/src/commands"
	"github.com/mjarkk/mini-exec/src/utils"
)

// ParseConf parses the .miniex file and runs the contents
func ParseConf() (commands.Command, error) {
	rawData, err := ioutil.ReadFile(".miniex")
	if err != nil {
		return commands.Command{}, err
	}

	data := strings.Split(string(rawData), "\n")
	for i, line := range data {
		data[i] = strings.Replace(line, "\r", "", -1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return commands.Command{}, err
	}

	toExec := []commands.Command{}
	finalExec := commands.Command{}
	commandCount := 0

	for num, line := range data {
		lineNum := num + 1

		if strings.HasPrefix(line, "#") || strings.Replace(strings.Replace(line, "\t", "", -1), " ", "", -1) == "" {
			continue
		}
		splitted := lineToArr(line)
		if len(splitted) == 0 {
			continue
		}

		if len(finalExec.Cmd) > 0 {
			return commands.Command{}, errors.New("Can't commands after FINAL, line: " + fmt.Sprint(lineNum))
		}

		isFinal := false
		firstArg := splitted[0]
		if firstArg == "FINAL" {
			if len(splitted) == 1 {
				return commands.Command{}, errors.New("`FINAL` needs more arguments, line: " + fmt.Sprint(lineNum))
			}
			splitted = splitted[1:]
			firstArg = splitted[0]
			isFinal = true
		}

		args := splitted[1:]

		if firstArg == "cd" {
			if isFinal {
				return commands.Command{}, errors.New("Final line can't be cd")
			}
			if len(args) == 0 {
				path, err := os.Getwd()
				if err != nil {
					return commands.Command{}, err
				}
				currentDir = path
				continue
			}

			commandCount++
			utils.Printf("---- Step %v Executing cd ----\n", commandCount)

			cdTo := path.Clean(args[0])
			if strings.HasPrefix(cdTo, "/") || strings.HasPrefix(cdTo, "\\") {
				currentDir = cdTo
				continue
			}
			currentDir = path.Join(currentDir, cdTo)
			continue
		}

		exec := commands.Command{
			Cmd:  firstArg,
			Args: args,
			Dir:  currentDir,
		}

		if isFinal {
			finalExec = exec
			continue
		}

		commandCount++
		utils.Printf("---- Step %v Executing %v ----\n", commandCount, exec.Cmd)
		exec.Prefix = true
		err = commands.Exec(exec)
		if err != nil {
			return commands.Command{}, err
		}

		toExec = append(toExec, exec)
	}

	return finalExec, nil
}

// lineToArr parses a shell line into a array that can be used for a exec
func lineToArr(in string) []string {
	arr := []string{""}
	splittedIn := strings.Split(in, "")

	goToNext := false
	isInsideQuotes := false
	quotesType := ""
	for _, part := range splittedIn {
		if (part == "\"" && (quotesType == "" || quotesType == "\"")) || (part == "'" && (quotesType == "" || quotesType == "'")) {
			isInsideQuotes = !isInsideQuotes
			if len(quotesType) == 1 {
				quotesType = ""
			} else {
				quotesType = part
			}
			part = ""
		}

		if part == " " && !isInsideQuotes {
			goToNext = true
			continue
		}

		if goToNext {
			arr = append(arr, "")
			goToNext = false
		}

		arr[len(arr)-1] += part
	}

	return arr
}
