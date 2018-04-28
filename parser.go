package golispy

import (
	"strings"
	"errors"
	"strconv"
	"fmt"
)

func Parse(input string) (Exp, error) {
	exp, _, err := ReadFromTokens(Tokenize(input))
	return exp, err
}

func Tokenize(input string) []string{
	input = strings.Replace(input,"(", " ( ", -1)
	input = strings.Replace(input,")", " ) ", -1)
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\t", "", -1)
	tokenized := strings.Split(input, " ")
	nonEmptyFunc := func(s string) bool {
		if strings.Trim(s, " ") == "" {
			return false
		}
		return true
	}
	tokenized = filter(tokenized, nonEmptyFunc)
	return tokenized
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func ReadFromTokens(tokens []string) (Exp, []string, error) {
	expression := Exp{}
	if len(tokens) == 0 {
		return expression, nil, errors.New("syntax error")
	}
	var token string
	token, tokens = tokens[0], tokens[1:]
	if token == "(" {
		list := List{}
		for tokens[0] != ")" {
			var subexp Exp
			var err error
			subexp, tokens, err = ReadFromTokens(tokens)
			if err != nil {
				return expression, nil, err
			}
			list = append(list, subexp)
		}
		tokens = tokens[1:] // Pop off ')'
		expression.list = &list
	} else if token == ")" {
		fmt.Print("ERROR")
		return expression, nil, errors.New("unexpected )")
	} else {
		atom := GetAtom(token)
		expression.atom = &atom
	}
	return expression, tokens,  nil
}

func GetAtom(token string) Atom {
	atom := Atom{}
	var err error

	var integer int64
	integer, err  = strconv.ParseInt(token, 10, 64)
	if err == nil {
		atom.integer = &integer
		return atom
	}

	var float float64
	float, err = strconv.ParseFloat(token, 64)
	if err == nil {
		atom.float = &float
		return atom
	}

	symbol := Symbol(token)
	atom.symbol = &symbol
	return atom
}
