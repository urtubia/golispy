package golispy

func DefaultEnv() Env {
	env := make(Env)
	env[">"] = greaterThanFunc
	env["<"] = lessThanFunc
	env["+"] = addFunc
	env["*"] = multFunc
	env["-"] = subtractFunc
	env["/"] = divideFunc
	env["list"] = listFunc
	env["car"] = carFunc
	env["cdr"] = cdrFunc
	env["begin"] = beginFunc

	return env
}

func Eval(x Exp, env Env) Exp {
	if x.IsSymbol() {
		return env[x.String()].(Exp)
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
	} else {
		list := *x.list
		proc := env[string(*list[0].atom.symbol)].(func([]Exp, Env)(Exp, error))
		var args List
		for _ , v := range list[1:] {
			args = append(args, Eval(v, env))
		}
		newExp, _ := proc(args, env)
		return newExp
	}
	return x
}
