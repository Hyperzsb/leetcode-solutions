type ParkingSystem struct {
	BigLen    int
	BigCap    int
	MediumLen int
	MediumCap int
	SmallLen  int
	SmallCap  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		BigLen:    0,
		BigCap:    big,
		MediumLen: 0,
		MediumCap: medium,
		SmallLen:  0,
		SmallCap:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if this.BigLen < this.BigCap {
			this.BigLen++
			return true
		} else {
			return false
		}
	} else if carType == 2 {
		if this.MediumLen < this.MediumCap {
			this.MediumLen++
			return true
		} else {
			return false
		}
	} else {
		if this.SmallLen < this.SmallCap {
			this.SmallLen++
			return true
		} else {
			return false
		}
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

