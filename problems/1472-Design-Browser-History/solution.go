type BrowserHistory struct {
	Idx  int
	URLs []string
}

func Constructor(homepage string) BrowserHistory {
	bh := BrowserHistory{Idx: 0, URLs: make([]string, 1)}
	bh.URLs[0] = homepage
	return bh
}

func (this *BrowserHistory) Visit(url string) {
	this.URLs = this.URLs[:this.Idx+1]
	this.URLs = append(this.URLs, url)
	this.Idx += 1
}

func (this *BrowserHistory) Back(steps int) string {
	if this.Idx >= steps {
		this.Idx -= steps
	} else {
		this.Idx = 0
	}

	return this.URLs[this.Idx]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.Idx+steps < len(this.URLs) {
		this.Idx += steps
	} else {
		this.Idx = len(this.URLs) - 1
	}

	return this.URLs[this.Idx]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

