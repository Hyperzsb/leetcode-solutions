type Day struct {
	day  int
	temp int
}

type Stack struct {
	s []Day
}

func (s *Stack) Push(day Day) {
	s.s = append(s.s, day)
}

func (s *Stack) Pop() Day {
	day := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return day
}

func (s *Stack) Top() Day {
	return s.s[len(s.s)-1]
}

func (s *Stack) Empty() bool {
	return len(s.s) == 0
}

func dailyTemperatures(temperatures []int) []int {
	dayStack := Stack{make([]Day, 0)}
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		if dayStack.Empty() {
			dayStack.Push(Day{i, temperatures[i]})
		} else if dayStack.Top().temp >= temperatures[i] {
			dayStack.Push(Day{i, temperatures[i]})
		} else {
			for !dayStack.Empty() && dayStack.Top().temp < temperatures[i] {
				day := dayStack.Pop()
				result[day.day] = i - day.day
			}
			dayStack.Push(Day{i, temperatures[i]})
		}
	}

	return result
}

