package main

import (
	"sgsh/shell"
)

// Enum for function exit codes
const (
	SUCCESS = iota
	FAILURE
)

func main() {
	shell.Loop()
}
