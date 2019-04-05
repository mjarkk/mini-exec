package app

import (
	"time"

	"github.com/mjarkk/mini-exec/src/commands"
	"github.com/mjarkk/mini-exec/src/utils"

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
				utils.Printf("BUILD STOPPED WITH ERROR:\n%v\n", err)
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
			utils.Println("---- Run FINAL command ----")
			toExec.Kill = kill
			err := commands.Exec(toExec)

			if !isKilled {
				if err != nil {
					utils.Println("FINAL COMMAND EXITED WITH CODE: ", err)
				} else {
					utils.Println("FINAL COMMAND EXITED WITHOUT ANY ERRORS")
				}
				<-kill
			} else {
				isKilled = false
				utils.Println("program killed, going to next release")
			}
		}
	}()
	go func() {
		for {
			time.Sleep(time.Minute * 2)
			status := GitPull()
			if !status {
				utils.Println("git pull: No update")
				continue
			}
			utils.Println("git pull: Got update")

			if isbuilding {
				newUpdate = true
				continue
			}

			newUpdateChan <- struct{}{}
		}
	}()
	return <-end
}
