package utils

import "fmt"

// Prefix is the application prefix for all logging related items
var Prefix = "[MINI-EXEC]"

// Println is the same as fmt.Println except this one adds a prefix
func Println(a ...interface{}) {
	fmt.Println(append([]interface{}{Prefix}, a...)...)
}

// Printf is the same as fmt.Printf except this one adds a prefix
func Printf(format string, args ...interface{}) {
	fmt.Printf(Prefix+" "+format, args...)
}
