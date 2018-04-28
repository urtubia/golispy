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

func TestMultilineParser(t *testing.T){
	golispyProgram1 := `
(begin
  (+ 1 2)
  (- 3 4)
  (define a 5)
  (define b 5)
  (define abc
    (lambda
		(a b c)
		(begin
			(+ 1 2)
			(+ 5 6)
			(+ a b c)
		)
    )
  )
  (abc 5 6 7)
)

`
	tokenized := golispy.Tokenize(golispyProgram1)
	expression, _, _ := golispy.ReadFromTokens(tokenized)
	expected := "( begin  ( +  1  2 )  ( -  3  4 )  ( define  a  5 )  ( define  b  5 )  ( define  abc  " +
		"( lambda  ( a  b  c )  ( begin  ( +  1  2 )  ( +  5  6 )  ( +  a  b  c ) ) ) )  ( abc  5  6  7 ) )"
	if expression.String() != expected {
		t.Error("Error parsing multiline program")
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
