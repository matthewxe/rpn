package main

import "testing"

func TestStack(t *testing.T) {
	t.Run("pushing and pulling a stack", func(t *testing.T) {
		stack := stack{}

		want := []int{3, 4, 6, 7}

		stack.Push(3)
		stack.Push(4)
		stack.Push(5)
		stack.Pull()
		stack.Push(6)
		stack.Push(7)
		stack.Push(8)
		stack.Pull()

		assertSlices(t, want, stack)
	})
	t.Run("pull an empty stack", func(t *testing.T) {
		stack := stack{}
		if _, ok := stack.Pull(); ok {
			t.Error("supposed to fail pulling an empty stack")
		}
	})
}

var hardEquation = []string{
	"3", "4", "*", "5", "+", "7", "2", "/", "1", "+", "-", "9", "*",
}

func TestEvalRPN(t *testing.T) {
	variations := []struct {
		name     string
		function func([]string) int
	}{
		{"EvalRPN", EvalRPN},
		{"EvalRPNFast", EvalRPNFast},
		{"EvalRPNFastLen", EvalRPNFastLen},
		{"EvalRPNFastPointers", EvalRPNFastPointers},
	}
	for _, v := range variations {
		t.Run(v.name+": 2 value addition", func(t *testing.T) {
			tokens := []string{"3", "4", "+"}
			want := 7
			got := v.function(tokens)
			assertEqual(t, want, got)
		})
		t.Run(v.name+": multi value operation", func(t *testing.T) {
			tokens := []string{"4", "13", "5", "+", "+"}
			want := 22
			got := v.function(tokens)
			assertEqual(t, want, got)
		})
		t.Run(v.name+": very hard equation", func(t *testing.T) {
			tokens := hardEquation
			want := 117
			got := v.function(tokens)
			assertEqual(t, want, got)
		})
	}
}

func assertSlices(t testing.TB, want, stack []int) {
	t.Helper()
	if len(want) != len(stack) {
		t.Error("stack and want dont have the same length")
	}

	for i := 0; i < len(want); i++ {
		if want[i] != stack[i] {
			t.Errorf("want %v, got %v", want, stack)
		}
	}
}

func assertEqual(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkEvalRPN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalRPN(hardEquation)
	}
}

func BenchmarkEvalRPNFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalRPNFast(hardEquation)
	}
}

func BenchmarkEvalRPNFastPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalRPNFastPointers(hardEquation)
	}
}

func BenchmarkEvalRPNFastLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalRPNFastLen(hardEquation)
	}
}
