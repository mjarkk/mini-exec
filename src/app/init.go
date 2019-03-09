package app

import (
	"fmt"
	"time"

	"github.com/mjarkk/mini-exec/src/commands"

	"github.com/mjarkk/mini-exec/src/files"
)

// Init is the start of the complte app
func Init() error {
	end := make(chan error)
	reRunFinal := make(chan commands.Command)

	isbuilding := false
	newUpdate := false
	newUpdateChan := make(chan struct{})

	GitPull()

	go func() {
		for {
			isbuilding = true
			toExec, err := files.ParseConf()
			if err != nil {
				fmt.Printf("[MINI-EXEC] BUILD STOPPED WITH ERROR:\n%v\n", err)
			} else {
				reRunFinal <- toExec
			}

			isbuilding = false
			if newUpdate {
				newUpdate = false
				continue
			}

			<-newUpdateChan
		}
	}()
	go func() {
		toExec := <-reRunFinal
		kill := make(chan struct{})
		isKilled := false
		go func() {
			for {
				toExec = <-reRunFinal
				kill <- struct{}{}
				isKilled = true
			}
		}()
		for {
			fmt.Println("[MINI-EXEC] ---- Run FINAL command ----")
			toExec.Kill = kill
			err := commands.Exec(toExec)

			if !isKilled {
				if err != nil {
					fmt.Println("[MINI-EXEC] FINAL COMMAND EXITED WITH CODE: ", err)
				} else {
					fmt.Println("[MINI-EXEC] FINAL COMMAND EXITED WITHOUT ANY ERRORS")
				}
				<-kill
			} else {
				isKilled = false
				fmt.Println("[MINI-EXEC] program killed, going to next release")
			}
		}
	}()
	go func() {
		for {
			time.Sleep(time.Minute * 1)
			status := GitPull()
			if !status {
				continue
			}

			if isbuilding {
				newUpdate = true
				continue
			}

			newUpdateChan <- struct{}{}
		}
	}()
	return <-end
}
