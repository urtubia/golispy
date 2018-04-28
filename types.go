package golispy

import (
	"fmt"
	"bytes"
)

/*
Symbol = str              # A Scheme Symbol is implemented as a Python str
Number = (int, float)     # A Scheme Number is implemented as a Python int or float
Atom   = (Symbol, Number) # A Scheme Atom is a Symbol or Number
List   = list             # A Scheme List is implemented as a Python list
Exp    = (Atom, List)     # A Scheme expression is an Atom or List
Env    = dict             # A Scheme environment (defined below) # is a mapping of {variable: value}
*/

type Symbol string

type Atom struct {
	symbol 	*Symbol
	integer *int64
	float 	*float64
}

type Callable interface {
	Call([]Exp, Env)(Exp, error)
}

type Exp struct {
	atom *Atom
	list *List
	callable Callable
}


func (e Exp) IsSymbol() bool{
	if e.atom != nil && e.atom.symbol != nil{
		return true
	}
	return false
}

func (e Exp) IsNumber() bool{
	if e.atom != nil && (e.atom.integer != nil || e.atom.float != nil) {
		return true
	}
	return false
}

func (e Exp) AsBool() bool {
	if e.atom != nil && e.atom.integer != nil && *e.atom.integer != 0 {
		return true
	}
	return false
}

func NewInt(i int64) Exp {
	atom := Atom{integer:&i}
	exp := Exp{atom:&atom}
	return exp
}

func (e Exp) IsEqual(other Exp) bool {
	if e.atom != nil && other.atom != nil {
		if e.atom.integer != nil && other.atom.integer != nil {
			if *e.atom.integer == *other.atom.integer {
				return true
			}
		}
		if e.atom.float != nil && other.atom.float != nil {
			if *e.atom.float == *other.atom.float {
				return true
			}
		}
		if e.atom.symbol != nil && other.atom.symbol != nil {
			if *e.atom.symbol == *other.atom.symbol {
				return true
			}
		}
	}
	if e.list != nil && other.list != nil {
		if len(*e.list) != len(*other.list) {
			return false
		}
		for i, v := range *e.list {
			if !v.IsEqual((*other.list)[i]) {
				return false
			}
		}
		return true
	}
	return false
}

type List []Exp

type Env map[string]Exp

func (e Exp) String() string {
	var buffer bytes.Buffer
	if e.list != nil {
		buffer.WriteString("(")
		for _, v := range *e.list {
			buffer.WriteString(" ")
			buffer.WriteString(v.String())
			buffer.WriteString(" ")
		}
		buffer.WriteString(")")
		return buffer.String()
	}
	return fmt.Sprint(e.atom)
}

func (a Atom) String() string {
	if a.integer != nil {
		return fmt.Sprintf("%d", *(a.integer))
	}
	if a.float != nil {
		return fmt.Sprintf("%f", *(a.float))
	}
	return fmt.Sprintf("%s", string(*(a.symbol)))
}

type Procedure struct {
	params List
	body Exp
	env Env
}

func (p Procedure) Call(args []Exp, env Env) (Exp, error) {
	return Exp{}, nil
}
