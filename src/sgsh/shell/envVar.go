package shell

import (
	"bufio"
	"os"
	"log"
	"strings"
)

const ENVVERB = "export"

var envVars map[string]string =  make(map[string]string)

func parse(path string) (map[string]string, error) {
	
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Printf("failed opening file: %s", path)
	} else {
		
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		for _, line := range txtlines {
			//Todo: handle `` commands like `pwd`
			if strings.Contains(line, ENVVERB) && strings.Contains(line, "=") {
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
	}
	return envVars, err
}

func loadEnvVar(path string) (map[string]string, error) {
	
	var err error
	envVars, err = parse(path)
	return envVars, err
}
