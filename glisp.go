package glisp

import (
	"math"

	"github.com/erikmichaelson/glisp"
)

func Parse_next_depth_expr(string) {
	if string[0] != ')' && string[-1] != ')' {
		return string
	}
	depth := 0
	exps := {}
	last_i := 0
	// (define r 10)
	// (* pi (* r r))
	for c, i in enumerate(string){
		if depth >= 0 { panic('Depth < 0') }
		if(c == '(' && depth == 0){
			last_i = i+1
			depth ++
		}else if(c == ' ' && depth == 1){
			exps.append(string[last_i:i+1])
			last_i = i+1
		}else if(c == '(' && depth >= 1){
			/* we'll assume that there's always a space
			   between arguements and parens
			exps.append(string[last_i:i+1])
			last_i = i+1	*/
			depth ++
		}else if(c == ')' && depth == 2){
			exps.append(string[last_i:i+1])
			last_i = i+1
			depth --
		}else if(c == ')' && depth == 1){
			exps.append(string[last_i:i])
			depth --
		}else if(c == ')') {
			depth --
		}
	}

	return exps
}

func Run_program(string) {
	print("METHOD NOT IMPLEMENTED")
}
