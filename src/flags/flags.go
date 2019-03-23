package flags

import "flag"

// OnlyCheck tells the application to only run the initial check function(s) and exit after that
var OnlyCheck = flag.Bool("errorCheck", false, "Only check for errors for errors and exit handy for in docker files before the CMD")

// Verbose prints out more information
var Verbose = flag.Bool("v", false, "verbose messages more information")

// AutoPull turns on / off auto timesouts
var AutoPull = flag.Bool("a", false, "git pull every view minutes to check if there are any updates")

// GitPullTimeout is the timeout between git pulls in minutes
var GitPullTimeout = flag.Uint("t", 2, "Time between git pull's in minutes")
