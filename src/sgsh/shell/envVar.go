package shell

import (
	"bufio"
	"os"
	"log"
	"strings"
)

var envVars map[string]string =  make(map[string]string)

func parse(path string) map[string]string {

	
	file, err := os.Open(path)
	if err != nil{
		log.Fatalf("failed opening file: %s", path)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	for _, line := range txtlines {
		//Todo: handle `` commands like `pwd`

		if strings.Contains(line, "define") && strings.Contains(line, "=") {
				fields := strings.Fields(line)
				for _ , field := range fields {
					if strings.Contains(field, "=") {
						tokens := strings.Split(field, "=")
						tokens[1] = strings.Trim(tokens[1], "\"")
						name, value := "$" + tokens[0], tokens[1]
						envVars[name] = value
						
				}	
			}
		
		}
	}

	return envVars
}

func loadEnvVar(path string) map[string]string {

	envVars = parse(path)
	return envVars
}
