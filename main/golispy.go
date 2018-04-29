package main

import (
	"golispy"
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	if len(os.Args) == 1 {
		golispy.StartRepl()
	} else {
		args := os.Args[1:]
		fileString, error := ioutil.ReadFile(args[0])
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file %s", fileString)
			return
		}
		env := golispy.DefaultEnv()
		parsedExp, _ := golispy.Parse(string(fileString))
		result := golispy.Eval(parsedExp, env)
		fmt.Println(result)
	}
}
