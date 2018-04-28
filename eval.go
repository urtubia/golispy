package golispy

import (
	"fmt"
	"os"
)

func DefaultEnv() Env {
	env := make(Env)
	env[">"] = Exp{callable: greaterThanCallable{}}
	env["<"] = Exp{callable: lessThanCallable{}}
	env["+"] = Exp{callable: addCallable{}}
	env["*"] = Exp{callable: multCallable{}}
	env["-"] = Exp{callable: subtractCallable{}}
	env["/"] = Exp{callable: divideCallable{}}
	env["list"] = Exp{callable: listCallable{}}
	env["car"] = Exp{callable: carCallable{}}
	env["cdr"] = Exp{callable: cdrCallable{}}
	env["begin"] = Exp{callable: beginCallable{}}

	return env
}

func Eval(x Exp, env Env) Exp {
	if x.IsSymbol() {
		return env[x.String()]
	} else if x.IsNumber() {
		return x
	} else if *(*x.list)[0].atom.symbol == "if" {
		list := *x.list
		test, conseq, alt := list[1], list[2], list[3]
		var result Exp
		if Eval(test, env).AsBool() {
			result = conseq
		}else{
			result = alt
		}
		return Eval(result, env)
	} else if *(*x.list)[0].atom.symbol == "define" {
		list := *x.list
		expSymbol, expValue := list[1], list[2]
		env[string(*expSymbol.atom.symbol)] = Eval(expValue, env)

	}else if *(*x.list)[0].atom.symbol == "lambda" {
		list := *x.list
		paramsExp, body := list[1], list[2]
		if paramsExp.list == nil {
			// TODO: Return error. Params here needs to be a list of symbols
			fmt.Fprint(os.Stderr, "Params to lambda needs to be a list of symbols")
		}
		proc := Procedure{params: *paramsExp.list, body: body, env: env}
		return Exp{callable:proc}
	} else {
		list := *x.list
		proc := env[string(*list[0].atom.symbol)].callable
		if proc != nil {
			var args List
			for _, v := range list[1:] {
				args = append(args, Eval(v, env))
			}
			newExp, _ := proc.Call(args, env)
			return newExp
		}
	}
	return x
}
