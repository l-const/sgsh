package shell

import (
	"fmt"
	"os"
)

const NUMBUILTINS = 4

type builtinFunc func([]string) int

var builtinArray [NUMBUILTINS]string

var mapBuiltinFn map[string]builtinFunc


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

	array := [NUMBUILTINS]string{"cd", "help", "exit", "load"}
	mapfunc := make(map[string]builtinFunc, NUMBUILTINS)
	builtinArray = array
	mapBuiltinFn = mapfunc
	mapBuiltinFn[array[0]] = cd
	mapBuiltinFn[array[1]] = help
	mapBuiltinFn[array[2]] = exit
	mapBuiltinFn[array[3]] = load
}

func help([]string) int {

	fmt.Printf("Konstantinos Lampropoulos's Simple Go Shell\n")
	for i := 0; i < NUMBUILTINS; i++ {
		fmt.Printf(" %s\n", builtinArray[i])
	}
	fmt.Printf("Use the man command for information on other programs\n")
	return 1
}

func exit([]string) int {

	return 0
}

func load(args [] string) int {

	// args[0] == load(command)
	for i, str := range args {
		if i != 0 {
			loadEnvVar(str)
		}
		
	}
	return 1
}
