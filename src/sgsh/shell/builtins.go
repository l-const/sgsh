package shell

import (
	"fmt"
	"os"
)

type builtinFunc func([]string) int

var builtinArray [3]string

var mapBuiltinFn map[string]builtinFunc

func numBuiltins() int {

	return len(builtinArray)
}

func cd(args []string) int {
	if len(args) < 2 {
		panic("Not enough arguments!")
	} else {
		err := os.Chmod(args[1], 0777)
		err = os.Chdir(args[1])
		if err != nil {
			panic(err)
		}
	}
	return 1
}

func initialize() {

	array := [3]string{"cd", "help", "exit"}
	mapfunc := make(map[string]builtinFunc, 3)
	builtinArray = array
	mapBuiltinFn = mapfunc
	mapBuiltinFn[array[0]] = cd
	mapBuiltinFn[array[1]] = help
	mapBuiltinFn[array[2]] = exit
}

func help([]string) int {

	fmt.Printf("Konstantinos Lampropoulos's Simple Go Shell\n")
	for i := 0; i < numBuiltins(); i++ {
		fmt.Printf("   %s\n", builtinArray[i])
	}
	fmt.Printf("Use the man command for information on other programs\n")
	return 1
}

func exit([]string) int {

	return 0
}
