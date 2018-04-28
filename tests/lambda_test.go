package tests

import "testing"

func TestLambda(t *testing.T){
	tests := map[string]string{
		"(begin (define addtwo (lambda (n) (+ n 2))) (addtwo 2))": "4",
		"(begin (define add (lambda (a b) (+ a b))) (add 2 5))": "7",
		"(begin (define fact (lambda (n) (if (< n 2) 1 (* n (fact (- n 1)))))) (fact 5))": "120"}
	runBasicTests(tests, t)
}


