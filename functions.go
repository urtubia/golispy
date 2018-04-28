package golispy

import (
	"errors"
)

type greaterThanCallable struct {}

func (g greaterThanCallable) Call(exps []Exp, env Env) (Exp, error) {
	exp := Exp{}
	var ret int64
	atom := Atom{integer: &ret}
	exp.atom = &atom
	if op1, op2 := Eval(exps[0], env), Eval(exps[1], env); len(exps) == 2 && op1.IsNumber() && op2.IsNumber() {
		if op1.atom.integer != nil && op2.atom.integer != nil {
			if *op1.atom.integer > *op2.atom.integer {
				ret = 1
			}
		}
		if op1.atom.float != nil && op2.atom.float != nil {
			if *op1.atom.float > *op2.atom.float {
				ret = 1
			}
		}
		return exp, nil
	}
	return exp, errors.New("operands to '>' must be 2 numbers")
}

type lessThanCallable struct {}

func (l lessThanCallable) Call(exps []Exp, env Env) (Exp, error) {
	exp := Exp{}
	var ret int64
	atom := Atom{integer: &ret}
	exp.atom = &atom
	if op1, op2 := Eval(exps[0], env), Eval(exps[1], env); len(exps) == 2 && op1.IsNumber() && op2.IsNumber() {
		if op1.atom.integer != nil && op2.atom.integer != nil {
			if *op1.atom.integer < *op2.atom.integer {
				ret = 1
			}
		}
		if op1.atom.float != nil && op2.atom.float != nil {
			if *op1.atom.float < *op2.atom.float {
				ret = 1
			}
		}
		return exp, nil
	}
	return exp, errors.New("operands to '>' must be 2 numbers")
}

type addCallable struct {}

func (a addCallable) Call(exps []Exp, env Env) (Exp, error) {
	opWithFloats := func(a float64,b float64) float64{
		return a + b
	}
	opWithInts := func(a int64,b int64) int64{
		return a + b
	}
	return numberOperandFunc(opWithFloats, opWithInts, exps, env)
}

type multCallable struct {}

func (m multCallable) Call(exps []Exp, env Env) (Exp, error) {
	opWithFloats := func(a float64,b float64) float64{
		return a * b
	}
	opWithInts := func(a int64,b int64) int64{
		return a * b
	}
	return numberOperandFunc(opWithFloats, opWithInts, exps, env)
}

type subtractCallable struct {}

func (m subtractCallable) Call(exps []Exp, env Env) (Exp, error) {
	opWithFloats := func(a float64,b float64) float64{
		return a - b
	}
	opWithInts := func(a int64,b int64) int64{
		return a - b
	}
	return numberOperandFunc(opWithFloats, opWithInts, exps, env)
}

type divideCallable struct {}

func (d divideCallable) Call(exps []Exp, env Env) (Exp, error) {
	opWithFloats := func(a float64,b float64) float64{
		return a / b
	}
	opWithInts := func(a int64,b int64) int64{
		return a / b
	}
	return numberOperandFunc(opWithFloats, opWithInts, exps, env)
}

func numberOperandFunc(opWithFloats func(float64,float64)float64,
	opWithInts func(int64, int64)int64, exps []Exp, env Env) (Exp, error){

	exp := Eval(exps[0].DeepAtomCopy(), env)
	if !exp.IsNumber() {
		return exp, errors.New("invalid operand type, must be number")
	}
	atom := exp.atom
	for _, op := range exps[1:] {
		evaled := Eval(op, env)
		if evaled.atom == nil || evaled.atom.symbol != nil {
			return exp, errors.New("invalid operand type")
		}

		if evaled.atom.integer != nil && atom.integer != nil {
			*atom.integer = opWithInts(*atom.integer, *evaled.atom.integer)
		}else if evaled.atom.float != nil && atom.integer != nil {
			var f float64
			atom.float = &f
			f = opWithFloats(*evaled.atom.float, float64(*atom.integer))
			atom.integer = nil
		}else if evaled.atom.integer != nil && atom.float != nil {
			*atom.float = opWithFloats(*atom.float, float64(*evaled.atom.integer))
		}else if evaled.atom.float != nil && atom.float != nil {
			*atom.float = opWithFloats(*atom.float, *evaled.atom.float)
		}
	}
	return exp, nil
}

type listCallable struct {}

func (l listCallable) Call(exps[]Exp, env Env) (Exp, error) {
	var list List
	for _, v := range exps {
		list = append(list, v)
	}
	return Exp{list: &list}, nil

}

type carCallable struct {}

func (l carCallable) Call(exps[]Exp, env Env) (Exp, error) {
	if len(exps) != 1 {
		return Exp{}, errors.New("car takes a single argument")
	}
	if *exps[0].list == nil {
		return Exp{}, errors.New("argument to car must be a list")
	}
	return (*exps[0].list)[0], nil
}

type cdrCallable struct {}

func (c cdrCallable) Call(exps[] Exp, env Env) (Exp, error) {
	if len(exps) != 1 {
		return Exp{}, errors.New("car takes a single argument")
	}
	if *exps[0].list == nil {
		return Exp{}, errors.New("argument to car must be a list")
	}

	argList := exps[0].list
	var newList List
	for _, v := range (*argList)[1:] {
		newList = append(newList, v)
	}

	return Exp{list:&newList}, nil
}

type beginCallable struct {}

func (b beginCallable) Call(exps[] Exp, env Env) (Exp, error) {
	return exps[len(exps)-1], nil
}