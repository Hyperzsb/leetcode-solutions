type TravelStats struct {
	From string
	To   string
	Cnt  int
	Sum  int
}

type CheckIn struct {
	From string
	Time int
}

type UndergroundSystem struct {
	TSMap map[string]*TravelStats
	CIMap map[int]*CheckIn
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		TSMap: make(map[string]*TravelStats),
		CIMap: make(map[int]*CheckIn),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.CIMap[id] = &CheckIn{
		From: stationName,
		Time: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	key := fmt.Sprintf("%s-%s", this.CIMap[id].From, stationName)

	if ts, ok := this.TSMap[key]; ok {
		ts.Cnt++
		ts.Sum += t - this.CIMap[id].Time
	} else {
		this.TSMap[key] = &TravelStats{
			From: this.CIMap[id].From,
			To:   stationName,
			Cnt:  1,
			Sum:  t - this.CIMap[id].Time,
		}
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := fmt.Sprintf("%s-%s", startStation, endStation)

	return float64(this.TSMap[key].Sum) / float64(this.TSMap[key].Cnt)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */

