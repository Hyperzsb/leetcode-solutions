func isValid(dist []int, hour float64, speed int) bool {
    totalHour := float64(0)
    for i := 0; i < len(dist)-1; i++ {
        h := float64(dist[i]) / float64(speed)
        if h > math.Trunc(h) {
            totalHour += float64(int(h) + 1)
        } else {
            totalHour += float64(int(h))
        }
    }

    totalHour += float64(dist[len(dist)-1]) / float64(speed)

    if totalHour > hour {
        return false
    } else {
        return true
    }
}

func minSpeedOnTime(dist []int, hour float64) int {
    totalDist := 0
    for i := range dist {
        totalDist += dist[i]
    }

    minSpeed, maxSpeed := 1, int(1e7)

    for minSpeed < maxSpeed {
        speed := (minSpeed + maxSpeed) / 2

        if isValid(dist, hour, speed) {
            maxSpeed = speed
        } else {
            minSpeed = speed + 1
        }
    }

    if isValid(dist, hour, minSpeed) {
        return minSpeed
    } else {
        return -1
    }
}

