package shell

import (
	"testing"
)

func TestProcessArgs(t *testing.T) {

	args := []string{"$NAME" , "$SURNAME", "$AGE"}
	expectedArgs := []string{"konstantinos", "Lampropoulos", "AGE"}
	vars := map[string]string {"$NAME" : "konstantinos", "$SURNAME" : "Lampropoulos"}
	resArgs, _ := ProcessArgs(args, vars)
	if len(resArgs) == len(expectedArgs) {
		for i := range resArgs {
			if resArgs[i] != expectedArgs[i] {
				t.Fatalf("Found error in processing env variabes!")
			}
		}
	
	}
}

func TestLaunch(*testing.T) {
	// Todo: Test builtins
	// Todo: Fix syscalls and test them
}