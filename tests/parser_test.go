package tests

import (
	"testing"
	"golispy"
	"strings"
)

func TestTokenizer(t *testing.T){
	program := "(begin (+ 1 2) (* 35 2 ) )"
	tokenized := golispy.Tokenize(program)
	expected := strings.Split("( begin ( + 1 2 ) ( * 35 2 ) )", " ")
	for i, v := range tokenized{
		if v != expected[i] {
			t.Error("%s != %s", v, expected[i])
		}
	}
}

func TestReadFromTokens(t *testing.T){
	program := "(begin (one 4) (* 32 34 (hello world)) blah other )"
	tokenized := golispy.Tokenize(program)
	expression, _, _ := golispy.ReadFromTokens(tokenized)
	expected := "( begin  ( one  4 )  ( *  32  34  ( hello  world ) )  blah  other )"
	result := expression.String()
	if expected != result {
		t.Error("Error creating AST expected %s got %s", expected, result)
	}
}
