package shell

import (
	"testing"
	"fmt"
)

func TestParsing(t *testing.T) {
	envVars, _ := parse("../.sgsh_profile")
	for name, val := range envVars {
		fmt.Printf("%s : %s\n", name, val)
	}
}