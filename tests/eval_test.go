package tests

import (
	"testing"
	"golispy"
)

func TestComparisonEval(t *testing.T){
	tests := map[string]string{
		"(> 1 2)": "0",
		"(< 1 2)": "1"}
	runBasicTests(tests, t)
}

func TestIfEval(t *testing.T){
	tests := map[string]string{
		"(if (< -1 0) 3 4)": "3",
		"(if (> 1 2) 3 4)": "4"}
	runBasicTests(tests, t)
}

func TestAdd(t *testing.T){
	tests := map[string]string{
		"(+ -0.0005 -143)": "-143.0005",
		"(+ 0.45 4)": "4.45",
		"(+ 0 -3 -10)": "-13",
		"(+ 10 -3 -10)": "-3",
		"(+ 10 -3)": "7",
		"(if (> (+ 1 11) -2) 1 2)":"1",
		"(+ 3 5)": "8"}
	runBasicTests(tests, t)
}

func TestMult(t *testing.T){
	tests := map[string]string{
		"(* 3 5)": "15",
		"(* 2 2 3)": "12",
		"(+ (- 4 (* 3 2) -6) 10 (/ 10 2))": "19",
		"(* (+ 1 1) (* 3 2))": "12"}
	runBasicTests(tests, t)
}

func TestCar(t *testing.T){
	tests := map[string]string{
		"(car (list 1 2 3))": "1",
		"(list 1 2 3))": "(1 2 3)"}
	runBasicTests(tests, t)
}


func TestCdr(t *testing.T){
	tests := map[string]string{
		"(cdr (list 1 2 3))": "(2 3)"}
	runBasicTests(tests, t)
}


func TestMisc(t *testing.T){
	tests := map[string]string{
		"(begin 1 (+ 3 2) 3 (list 4 4))": "(4 4)",
		"(begin (define d 10) (define p 100) (+ p d))": "110",
		"(begin (define a 10) (define b 100) (list (+ a b) b))": "(110 100)",
		"(begin 1 2 3 4 5)": "5"}
	runBasicTests(tests, t)
}

func runBasicTests(tests map[string]string, t *testing.T){
	for k, v := range tests {
		env := golispy.DefaultEnv()
		testExpression, _ := golispy.Parse(k)
		testExpression = golispy.Eval(testExpression, env)
		expected, _ := golispy.Parse(v)
		if !testExpression.IsEqual(expected) {
			t.Error(testExpression.String(), " not equal to " ,expected.String(), " when testing ", k)
		}
	}
}

