package flags

import "flag"

// OnlyCheck tells the application to only run the initial check function(s) and exit after that
var OnlyCheck = flag.Bool("errorCheck", false, "Only check for errors for errors and exit")

// Verbose prints out more information
var Verbose = flag.Bool("v", false, "verbose messages (more information)")
