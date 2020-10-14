package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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

// Execute : Run the parsed command.
func Execute(args []string) int {
	if len(args) == 0 {
		// Empty command
		return 1
	}
	initialize()
	for i := 0; i < numBuiltins(); i++ {
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
		var procAttr os.ProcAttr
		procAttr.Dir = "."
		procAttr.Env = os.Environ()
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		proc, err := os.StartProcess(args[0], args[1:], &procAttr)
		if err == nil {
			proc.Wait()
			return 1
		}
		panic(err)
	}
	// Todo: add environmetal variables and pass them
	// Todo: in the execution
	cmd := exec.Command(args[0], args[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(stdout))
	return 1
}

// Loop sdsdsds.
func Loop() {

	status := 1
	for status != 0 {
		fmt.Printf("[$sgsh] > ")
		line := ReadLine()
		args := SplitLine(string(line))
		status = Execute(args)
	}
}
