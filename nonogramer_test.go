package main

import (
	"testing"
)

func testSolver(input string, expected string, t *testing.T) {
	if solve(input) != expected {
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	testSolver("\n", "\n", t)
}

func TestFilled1x1(t *testing.T) {
	testSolver(" 1\n1", "O\n", t)
}

func TestEmpty1x1(t *testing.T) {
	testSolver(" 0\n0", ".\n", t)
}

func TestFilled2x2(t *testing.T) {
	testSolver(" 22\n2\n2", "OO\nOO\n", t)
}
