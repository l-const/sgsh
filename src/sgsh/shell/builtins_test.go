package shell

import (
	"testing"
)

func TestLoad(t *testing.T) {
	args := []string{"load", "../.sgshrc", "../.sgsh_profile"}
	_, err := load(args)
	if err != nil {
		t.Fatalf("Error in load()")
	}
}

func TestChdir(*testing.T) {
	
}
