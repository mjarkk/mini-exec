package server

import (
	"github.com/gin-gonic/gin"
)

// StateT is the State type
type StateT struct {
	Commands []Step
}

// Step is the output of a command
type Step struct {
	Command string   // the actual command
	Output  []string // the text output of the command
	ExitErr bool     // did the command exit with an error
}

// State is the web UI it's state
var State = StateT{
	Commands: []Step{},
}

// SendState sends the current state
func sendState(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
		"state":  State,
	})
}
