package main

import (
	"fmt"
	"unittest"
)
// TODO: translate this
from lis import *

//class TestHelpers(unittest.TestCase):
func TestParser(t *testing.T) {
	// looks like an efficient way to do
	// subtests
	var tests = []struct {
		input string
		expects []string
	}{
		{"(define r 10)", {"define","r","10"}},
		{"(* pi (* r r))", {"*","pi","(* r r)"}},
		{"(* a (- b c)(+ d e (* f g)))", {"*", "a","(- b c)", "(+ d e (* f g))"}}
	}

	for _, test := range tests {
		testname := input

		t.Run(testname, func(t *testing.T) {
			ans := parse_next_depth_expr(input)
			if ans != test.expects {
				t.Errorf("got %s, want %s", ans, test.expects)
			}
		})
	}
	/*
	func test_parser_1(self):
		program = "(define r 10)"
		exps = parse_next_depth_expr(program)
		self.assertEqual(exps, ['define','r','10'])
		self.assertEqual(len(exps), 3)

	def test_parser_2(self):
		program = "(* pi (* r r))"
		exps = parse_next_depth_expr(program)
		self.assertEqual(exps, ['*','pi','(* r r)'])
		self.assertEqual(len(exps), 3)

	def test_parser_3(self):
		program = "(* a (- b c)(+ d e (* f g)))"
		exps = parse_next_depth_expr(program)
		self.assertEqual(exps, ['*', 'a','(- b c)', '(+ d e (* f g))'])
		self.assertEqual(len(exps), 4)
	*/
}


class TestLiterals(unittest.TestCase):
	def test_constant(self):
		result = run_program("12")
		self.assertEqual(result, 12,
			f"expected 12 but recieved {result}")
	
class TestBasicLists(unittest.TestCase):
	def test_mult(self):
		result = run_program("(* 3 2 1 14)")
		self.assertEqual(result, 84,
			f"expected 84 but recieved {result}")
	def test_sqrt(self):
		result = run_program("(sqrt (* 2 8))")
		self.assertEqual(result, 4.0,
			f"expected 4 but recieved {result}")
	def test_define_and_reference(self):
		result = run_program("(define r 10)\nr")
		self.assertEqual(result, 10,
			f"expected 10 but recieved {result}")
	def test_circle_area(self):
		program = "(define r 10)\n(* pi (* r r))"
		result = run_program(program)
		self.assertEqual(result, math.pi * 100,
			f"expected 314.15 but recieved {result}")
	def test_nested_mult(self):
		program = "(* 2 (- 3 4)(+ 9 -2 (* 5 1)))"
		result = run_program(program)
		# 2 * -1 * (7 + 5) = -24
		self.assertEqual(result, -24,
			f"expected 24 but recieved {result}")

class TestAdvanced(unittest.TestCase):
	def test_function_gen(self):
		program = \
		"(if (> (val x) 0)\n(fn (+ (aref A i) (* 3 i))\n(quote (one two)))"
		result = run_program(program)

if __name__ == '__main__':
	unittest.main()
