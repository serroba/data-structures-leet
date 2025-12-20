package training

// [73, 74, 75, 71, 69, 72, 76, 73]
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	s := stackInt{}
	for i := 0; i < len(temperatures); i++ {
		for !s.IsEmpty() && temperatures[i] > temperatures[s.Peek()] {
			j := s.Pop()
			res[j] = i - j
		}
		s.Push(i)
	}
	return res
}

type stackInt []int

func (s *stackInt) Len() int {
	return len(*s)
}

func (s *stackInt) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stackInt) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stackInt) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *stackInt) Push(i int) {
	*s = append(*s, i)
}
