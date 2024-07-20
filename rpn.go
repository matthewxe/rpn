package main

import "strconv"

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pull() (val int, ok bool) {
	endix := len(*s) - 1
	if endix < 0 {
		return 0, false
	}
	val = (*s)[endix]
	*s = (*s)[:endix]
	return val, true
}

func (s *stack) Eval(operator string) {
	b, _ := s.Pull()
	a, _ := s.Pull()
	switch operator {
	case "+":
		s.Push(a + b)
	case "-":
		s.Push(a - b)
	case "*":
		s.Push(a * b)
	case "/":
		s.Push(a / b)
	}
}

func EvalRPN(tokens []string) int {
	stack := stack{}

	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err != nil {
			stack.Eval(v)
		} else {
			stack.Push(val)
		}
	}

	out, ok := stack.Pull()
	if !ok {
		return -1
	}
	return out
}

func EvalRPNFast(tokens []string) int {
	var s []int
	sIndex := 0
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			sIndex--
			switch token {
			case "+":
				s[sIndex-1] += s[sIndex]
			case "-":
				s[sIndex-1] -= s[sIndex]
			case "*":
				s[sIndex-1] *= s[sIndex]
			case "/":
				s[sIndex-1] /= s[sIndex]
			}
			s = s[:sIndex]
		} else {
			s = append(s, num)
			sIndex++
		}
	}
	return s[0]
}

func EvalRPNFastLen(tokens []string) int {
	var s []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			sIndex := len(s) - 1
			switch token {
			case "+":
				s[sIndex-1] += s[sIndex]
			case "-":
				s[sIndex-1] -= s[sIndex]
			case "*":
				s[sIndex-1] *= s[sIndex]
			case "/":
				s[sIndex-1] /= s[sIndex]
			}
			s = s[:sIndex]
		} else {
			s = append(s, num)
		}
	}
	return s[0]
}

func EvalRPNFastPointers(tokens []string) int {
	var s []int
	sIndex := 0
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			sIndex--
			b := s[sIndex]
			a := &(s[sIndex-1])
			s = s[:sIndex]
			switch token {
			case "+":
				*a = *a + b
			case "-":
				*a = *a - b
			case "*":
				*a = *a * b
			case "/":
				*a = *a / b
			}
		} else {
			s = append(s, num)
			sIndex++
		}
	}
	return s[0]
}
