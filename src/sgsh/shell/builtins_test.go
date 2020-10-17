package shell

import (
	"testing"
)



func TestLoad(t *testing.T) {
	args := []string{"load", "../.sgshrc", "../.sgsh_profile"}
	load(args)
}