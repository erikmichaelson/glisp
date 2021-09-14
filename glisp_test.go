package glisp_test

import (
	"math"
	"testing"

	"github.com/erikmichaelson/glisp"
)
// TODO: translate this
//from lis import *

//class TestHelpers(unittest.TestCase):
func TestParser(t *testing.T) {
	// looks like an efficient way to do
	// subtests
	var tests = []struct {
		input string
		expects []string
	}{
		{"(define r 10)", []string{"define","r","10"}},
		{"(* pi (* r r))", []string{"*","pi","(* r r)"}},
		{"(* a (- b c)(+ d e (* f g)))", []string{"*", "a","(- b c)", "(+ d e (* f g))"}}}

	for _, test := range tests {
		testname := test.input

		t.Run(testname, func(t *testing.T) {
			ans := Parse_next_depth_expr(test.input)
			if ans != test.expects {
				t.Errorf("got %s, want %s", ans, test.expects)
			}
		})
	}
}


func TestLiterals(t *testing.T) {
	t.Run("Literal", func(t *testing.T) {
				ans := Run_program("12")
				if ans != 12 {
					t.Errorf("got %s, want %s", ans, 12)
				}
			})
}

func TestBasicLists(t *testing.T) {
	var tests = []struct {
			input string
			expects float64
		}{
			{"(* 3 2 1 14)",			84.0},
			{"(sqrt (* 2 8))",			4.0},
			{"(define r 10)\nr",			10.0},
			{"(define r 10)\n(* pi (* r r))",	math.Pi * 100},
			{"(* 2 (- 3 4)(+ 9 -2 (* 5 1)))",	-24.0}}

		for _, test := range tests {
			testname := test.input

			t.Run(testname, func(t *testing.T) {
				ans := Run_program(test.input)
				if ans != test.expects {
					t.Errorf("got %s, want %s", ans, test.expects)
				}
			})
		}
}

/*
need to figure out how to store and represent a function stored in execution to test this
TODO

func TestAdvanced(t *testing.T) {
	t.Run("Advanced", func(t *testing.T) {
		ans := Run_program("(if (> (val x) 0)\n(fn (+ (aref A i) (* 3 i))\n(quote (one two)))")
		if ans != test.expects {
			t.Errorf("got %s, want %s", ans, test.expects)
		}
	})
}	*/

