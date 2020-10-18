package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"log"
	"errors"

)

func must (err error) {
	if err != nil {
		panic(err)
	}
}

// ReadLine : Read the command from standard input.
func readLine() []byte {

	buffer := make([]byte, 0, 100)
	reader := bufio.NewReader(os.Stdin)

	for {
		char, error := reader.ReadByte()
		if error != nil {
			if error == io.EOF {
				return buffer
			}
			panic("Error reading a byte!")
		}

		if char == '\n' {
			return buffer
		}
		buffer = append(buffer, char)
	}
}

// SplitLine : Separate the command string into a program and arguments.
func splitLine(line string) []string {

	words := strings.Fields(line)
	return words
}


func processArgs(args [] string, vars map[string]string) ([]string,  error) {

	// args := ["$NAME", "cd", ... ]
	// vars := ["$NAME": "KONSTANSTINOS"]
	var err error
	for i, v := range args {
		if strings.Contains(v, "$") {
			if _ , ok := vars[args[i]]; ok {
				args[i] = vars[args[i]]
			} else {
				log.Print(args[i] + " not defined!")
				err = errors.New(args[i] + " not defined!")
			}
		}
	}

	return args, err 
}


func insertSlice(a []string, index int, value string) []string {
    if len(a) == index || index == -1 { // nil or empty slice or after last element
        return append(a, value)
    }
    a = append(a[:index+1], a[index:]...) // index < len(a)
    a[index] = value
    return a
}


func pipe(first [] string, rest [] string) (int, error) {
	// 1. Execute first
	// 2. take output 
	// 3. check rest for pipe 
	// a. if has => find position (i) and insert output string  in i-1
	// 	a1. set first equal to rest[:i]
	//  a2. set rest equal to rest[i:]
	//  a3. call recursively pipe(first, equal) 
	// b. if doesn't have pipe symbol 
	//   b1. append output to rest.
	//   b2. call and return  checkTypecommand(rest)
	var newrest [] string
	_, err := exec.LookPath(first[0])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unkown command")
	} else {
		if len(first) > 1 {
			cmd := exec.Command(first[0], first[1:]...)
			stdout, err := cmd.Output()
			must(err)
			output := string(stdout)
			for i, v := range rest {
				if strings.Contains(v, "|") {
					newrest = insertSlice(rest, i, output)
					first = newrest[:i+1]
					rest = newrest[i+2:]
					pipe(first, rest)
				}
			}
			newrest = insertSlice(rest, -1, output)

		} else {
			fmt.Println("Not enough arguments")
		}
	}
	return checkTypeCommand(newrest)
}


func checkTypeCommand(args [] string) (int, error) {

	//builtins
	for i := 0; i < NUMBUILTINS; i++ {
		if args[0] == builtinArray[i] {
			return mapBuiltinFn[builtinArray[i]](args)
		}
	}
	// no builtin command
	return launch(args)
}

// execute : Run the parsed command.
func execute(args []string) (int, error) {

	var err error
	if len(args) == 0 {
		// Empty command
		return 1, err
	}
	// check for chained commands => &||&&
	for i, v := range args {
		if strings.Contains(v, "&") {
			// multiple commands
			j, _ := execute(args[:i])
			i, _ := execute(args[i+1:])
			return i & j, nil
		}

	}
	// check for pipes -> |
	for i, v := range args {
		if strings.Contains(v, "|") {
			return pipe(args[:i], args[i+1:])
		}
	}

	return checkTypeCommand(args)
}

// launch : Run the parsed command.
func launch(args []string) (int, error) {

	_, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unkown command")
	} else {
		if len(args) > 1 {
			cmd := exec.Command(args[0], args[1:]...)
			stdout, err := cmd.Output()
			must(err)
			fmt.Print(string(stdout))
		} else {
			fmt.Println("Not enough arguments")
		}
	}
	return 1, err
}

// Loop execution.
func Loop() {
	var err error
	status := 1
	initialize()
	vars, _ := loadEnvVar(".sgsh_profile")
	for status != 0 {
		fmt.Printf("[$sgsh] > ")
		line := readLine()
		args := splitLine(string(line))
		args, err = processArgs(args, vars)
		if err != nil {
			continue
		}
		status, err = execute(args)
	}
}
