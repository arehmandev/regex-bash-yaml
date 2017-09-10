package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	filepath   = "variables.sh"
	globalvars = "globalvars.yaml"
)

func main() {
	createyaml(globalvars)
	bashtoyaml(filepath, globalvars)
	bashtoyaml(filepath, globalvars) // can add more than one bash file and append it to the yaml
}

func createyaml(destination string) {
	os.Create(destination)

}

func bashtoyaml(source, destination string) {
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// os.Create(destinationx)
	for scanner.Scan() {
		// each line's test as a string
		stringline := scanner.Text()

		// change bash variables to yaml by removing first equals and replacing with colon space

		pat := regexp.MustCompile("^(.*?)=(.*)$")
		repl := "${1}: $2"
		output := pat.ReplaceAllString(stringline, repl)

		fmt.Println(output)
		// write to file

		f, err := os.OpenFile(destination, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		fmt.Fprintf(f, "%s \n", output)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
