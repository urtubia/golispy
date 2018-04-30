# Dizzy Golispy

Dizzy Golispy is a simple LISP interpreter in Go, based on Peter Norvig's [lispy](http://norvig.com/lispy.html) language.
It is currently under development, but its goals are:
* Make it easy to embed
* Keep base language simple and minimal

## Running

### Run unit tests
`go test ./tests`

### Build executable
`go build main/golispy`

### Usage
Start an interactive repl:
`./golispy`

Interpret a file with a golispy program:
`./golispy program.scm`

