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

// ReadLine : Read the command from standard input.
func ReadLine() []byte {

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
func SplitLine(line string) []string {

	words := strings.Fields(line)
	return words
}

func ProcessArgs(args [] string,vars map[string]string) ([]string,  error) {

	// args := ["$NAME", "cd", ... ]
	/// vars := ["$NAME": "KONSTANSTINOS"]
	var err error
	for i, v := range args {
		if strings.Contains(v, "$") {
			if _ , ok := vars[args[i]]; ok {
				args[i] = vars[args[i]]
			} else {
				log.Print(args[i] + " not defined!")
				err = errors.New(args[i] + " not defined!")
			}
			fmt.Println(args[i])
		}
	}

	return args, err 
}

// Execute : Run the parsed command.
func Execute(args []string) int {

	if len(args) == 0 {
		// Empty command
		return 1
	}
	initialize()
	for i := 0; i < NUMBUILTINS; i++ {
		if args[0] == builtinArray[i] {
			return mapBuiltinFn[builtinArray[i]](args)
		}
	}
	return Launch(args)
}

// Launch : Run the parsed command.
func Launch(args []string) int {

	_, err := exec.LookPath(args[0])
	if err != nil {
		
		fmt.Println(err)
		fmt.Println("Unkown command or implemented by syscall!")
	}

	cmd := exec.Command(args[0], args[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(stdout))
	return 1
}

// Loop execution.
func Loop() {
	var err error
	status := 1
	for status != 0 {
		fmt.Printf("[$sgsh] > ")
		line := ReadLine()
		args := SplitLine(string(line))
		vars := loadEnvVar(".sgsh_profile")
		args, err = ProcessArgs(args, vars)
		if err != nil {
			continue
		}
		status = Execute(args)
	}
}
