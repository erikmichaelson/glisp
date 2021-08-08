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
}


func TestLiterals(t *testing.T):
	t.Run(testname, func(t *testing.T) {
				ans := run_program("12")
				if ans != 12 {
					t.Errorf("got %s, want %s", ans, test.expects)
				}
			})
	
func TestBasicLists(t *testing.T):
	var tests = []struct {
			input string
			expects []string
		}{
			{"(* 3 2 1 14)", 			84},
			{"(sqrt (* 2 8))",			4.0},
			{"(define r 10)\nr", 			10},
			{"(define r 10)\n(* pi (* r r))", 	math.pi * 100},
			{"(* 2 (- 3 4)(+ 9 -2 (* 5 1)))", 	-24}
		}

		for _, test := range tests {
			testname := input

			t.Run(testname, func(t *testing.T) {
				ans := run_program(input)
				if ans != test.expects {
					t.Errorf("got %s, want %s", ans, test.expects)
				}
			})
		}

func TestAdvanced(t *testing.T):
	t.Run(testname, func(t *testing.T) {
		ans := run_program("(if (> (val x) 0)\n(fn (+ (aref A i) (* 3 i))\n(quote (one two)))")
		if ans != test.expects {
			t.Errorf("got %s, want %s", ans, text.expects)
		}
	})

if __name__ == '__main__':
	unittest.main()
