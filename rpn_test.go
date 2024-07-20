package main

import "testing"

var hardEquation = []string{
	"3", "4", "*", "5", "+", "7", "2", "/", "1", "+", "-", "9", "*",
}

func TestEvalRPN(t *testing.T) {
	t.Run("2 value addition", func(t *testing.T) {
		tokens := []string{"3", "4", "+"}
		want := 7
		got := EvalRPN(tokens)
		assertEqual(t, want, got)
	})
	t.Run("multi value operation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPN(tokens)
		assertEqual(t, want, got)
	})
	t.Run("very hard equation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPN(tokens)
		assertEqual(t, want, got)
	})
}

func TestEvalRPNFast(t *testing.T) {
	t.Run("2 value addition", func(t *testing.T) {
		tokens := []string{"3", "4", "+"}
		want := 7
		got := EvalRPNFast(tokens)
		assertEqual(t, want, got)
	})
	t.Run("multi value operation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPNFast(tokens)
		assertEqual(t, want, got)
	})
	t.Run("very hard equation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPNFast(tokens)
		assertEqual(t, want, got)
	})
}

func TestEvalRPNFastPointers(t *testing.T) {
	t.Run("2 value addition", func(t *testing.T) {
		tokens := []string{"3", "4", "+"}
		want := 7
		got := EvalRPNFastPointers(tokens)
		assertEqual(t, want, got)
	})
	t.Run("multi value operation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPNFastPointers(tokens)
		assertEqual(t, want, got)
	})
	t.Run("very hard equation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPNFastPointers(tokens)
		assertEqual(t, want, got)
	})
}

func TestEvalRPNFastLen(t *testing.T) {
	t.Run("2 value addition", func(t *testing.T) {
		tokens := []string{"3", "4", "+"}
		want := 7
		got := EvalRPNFastLen(tokens)
		assertEqual(t, want, got)
	})
	t.Run("multi value operation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPNFastLen(tokens)
		assertEqual(t, want, got)
	})
	t.Run("very hard equation", func(t *testing.T) {
		tokens := []string{"4", "13", "5", "/", "+"}
		want := 6
		got := EvalRPNFastLen(tokens)
		assertEqual(t, want, got)
	})
}

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

func BenchmarkEvalRPNFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalRPNFast(hardEquation)
	}
}
