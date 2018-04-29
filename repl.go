package golispy

import (
	"bufio"
	"os"
	"fmt"
)

func StartRepl() {
	reader := bufio.NewReader(os.Stdin)
	env := DefaultEnv()
	for {
		fmt.Print("golispy > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		parsedExp, _ := Parse(line)
		result := Eval(parsedExp, env)
		fmt.Println(result)
	}
}
