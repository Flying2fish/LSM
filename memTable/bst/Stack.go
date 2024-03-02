package bst

type Stack struct {
	stack  []*treeNode
	length int
}

func (s *Stack) Push(t *treeNode) {
	s.stack = append(s.stack, t)
	s.length++
}

func (s *Stack) Pop() (*treeNode, bool) {
	if s.length == 0 {
		return nil, false
	}
	v := s.stack[s.length-1]
	s.stack = s.stack[:s.length-1]
	s.length--
	return v, true
}

func (s *Stack) Len() int {
	return s.length
}
