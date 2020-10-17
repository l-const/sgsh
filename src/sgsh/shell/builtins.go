package shell

import (
	"fmt"
	"syscall"
	"log"
	"errors"
)

const NUMBUILTINS = 4

type builtinFunc func([]string) (int, error)

var builtinArray [NUMBUILTINS]string

var mapBuiltinFn map[string]builtinFunc


func cd(args []string) (int, error) {

	var err error
	if len(args) < 2 {
		log.Printf("Not enough arguments!")
		err = errors.New("Not enough arguments!")
	} else {
		err = syscall.Chdir(args[1])
		if err != nil {
			fmt.Println("error chdir")
		}
		fmt.Printf("SURVIVED CHDIR!")
	}
	return 1, err
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

func help([]string) (int, error) {
	
	var err error
	fmt.Printf("Konstantinos Lampropoulos's Simple Go Shell\n")
	for i := 0; i < NUMBUILTINS; i++ {
		fmt.Printf(" %s\n", builtinArray[i])
	}
	fmt.Printf("Use the man command for information on other programs\n")
	return 1, err
}

func exit([]string) (int, error) {

	var err error
	return 0, err
}

func load(args [] string) (int, error) {

	var err error
	// args[0] == load(command)
	for i, str := range args {
		if i != 0 {
			_, err = loadEnvVar(str)
		}
		
	}
	return 1, err
}
