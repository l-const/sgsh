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

func TestExecute(t *testing.T) {
	tests := [][]string{{"load", "../.sgsh_profile"}, {"help"}, {"exit"}}
	for i, v  := range tests {
		_, err := Execute(v)
		if err != nil {
			t.Fatalf("Error in %s test",tests[i][0])
		}
	}
	
}