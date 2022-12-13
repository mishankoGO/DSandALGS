package tasks

type Stack struct {
	data []string
}

func (s *Stack) push(x string) {
	s.data = append(s.data, x)
}

func (s *Stack) pop() string {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) peek() string {
	if len(s.data) == 0 {
		return ""
	}
	val := s.data[len(s.data)-1]
	return val
}

var Brackets = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func IsCorrectBracketsSeq(brackets string) bool {
	var brStack Stack
	for _, elem := range brackets {
		if bracket, ok := Brackets[string(elem)]; ok {
			curr := brStack.peek()
			if curr != bracket {
				return false
			} else {
				brStack.pop()
			}
		} else {
			brStack.push(string(elem))
		}
	}
	if len(brStack.data) != 0 {
		return false
	}
	return true
}
