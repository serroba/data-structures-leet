package training

// [73, 74, 75, 71, 69, 72, 76, 73]
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	s := stack{}
	for i := 0; i < len(temperatures); i++ {
		for !s.IsEmpty() && temperatures[i] > temperatures[s.Peek()] {
			j := s.Pop()
			res[j] = i - j
		}
		s.Push(i)
	}
	return res
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) Push(i int) {
	*s = append(*s, i)
}
