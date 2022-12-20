func canVisitAllRooms(rooms [][]int) bool {
	roomsAccessable := make([]bool, len(rooms))
	roomsAccessable[0] = true

	roomsAvailable := make([]int, 1)
	roomsAvailable[0] = 0
	for i := 0; i < len(roomsAvailable); i++ {
		for j := 0; j < len(rooms[roomsAvailable[i]]); j++ {
			if !roomsAccessable[rooms[roomsAvailable[i]][j]] {
				roomsAccessable[rooms[roomsAvailable[i]][j]] = true
				roomsAvailable = append(roomsAvailable, rooms[roomsAvailable[i]][j])
			}
		}
	}

	for _, val := range roomsAccessable {
		if !val {
			return false
		}
	}

	return true
}

