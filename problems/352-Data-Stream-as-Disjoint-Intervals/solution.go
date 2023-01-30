type SummaryRanges struct {
	GetRange map[int]*Range
	Ranges   []*Range
}

type Range struct {
	Start int
	End   int
}

func Constructor() SummaryRanges {
	return SummaryRanges{GetRange: make(map[int]*Range),
		Ranges: make([]*Range, 0, 1)}
}

func (this *SummaryRanges) AddNum(value int) {
	if r := this.GetRange[value]; r != nil {
		return
	}

	rl := this.GetRange[value-1]
	rr := this.GetRange[value+1]
	if rl == nil && rr == nil {
		r := Range{Start: value, End: value}
		this.GetRange[value] = &r
		this.Ranges = append(this.Ranges, &r)
	} else if rl != nil && rr == nil {
		this.GetRange[value] = rl
		(*rl).End = value
	} else if rl == nil && rr != nil {
		this.GetRange[value] = rr
		(*rr).Start = value
	} else {
		for i := value; i <= (*rr).End; i++ {
			this.GetRange[i] = rl
		}

		(*rl).End = (*rr).End

		for i := 0; i < len(this.Ranges); i++ {
			if this.Ranges[i] == rr {
				this.Ranges = append(this.Ranges[:i], this.Ranges[i+1:]...)
				break
			}
		}
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	sort.Slice(this.Ranges, func(a, b int) bool {
		return (*this.Ranges[a]).Start < (*this.Ranges[b]).Start
	})

	result := make([][]int, len(this.Ranges))
	for i, r := range this.Ranges {
		result[i] = make([]int, 2)
		result[i][0] = (*r).Start
		result[i][1] = (*r).End
	}

	return result
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */

