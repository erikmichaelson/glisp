package glisp

/*
import (
	"math"
)	*/

func parse_next_depth_expr(str string) []string {
	if len(str) < 1 { panic("Empty string passed to parser") }

	if rune(str[0]) != ')' && rune(str[len(str)-1:][0]) != ')' {
		return []string {str}
	}

	print("it got HERE\n")
	depth := 0
	var exps []string
	last_i := 0
	// (define r 10)
	// (* pi (* r r))
	for c, i := range str {
		if depth < 0 { panic("Depth < 0") }
		if(c == '(' && depth == 0){
			last_i = int(i+1)
			depth ++
		}else if(c == ' ' && depth == 1){
			exps = append(exps, str[last_i:int(i+1)])
			last_i = int(i+1)
		}else if(c == '(' && depth >= 1){
			/* we'll assume that there's always a space
			   between arguements and parens
			exps.append(str[last_i:int(i+1)])
			last_i = int(i+1)	*/
			depth ++
		}else if(c == ')' && depth == 2){
			exps = append(exps, str[last_i:int(i+1)])
			last_i = int(i+1)
			depth --
		}else if(c == ')' && depth == 1){
			exps = append(exps, str[last_i:i])
			depth --
		}else if(c == ')') {
			depth --
		}
	}

	return exps
}

func run_program(str string) string {
	return "METHOD NOT IMPLEMENTED"
}
