func findOrder(numCourses int, prerequisites [][]int) []int {
	degrees := make([]int, numCourses)
	prereqs := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		prereqs[i] = make([]int, 0)
	}

	for _, val := range prerequisites {
		degrees[val[0]]++
		prereqs[val[1]] = append(prereqs[val[1]], val[0])
	}

	coursesTaken := make([]bool, numCourses)
	coursesAvailable := make([]int, 0)
	coursesList := make([]int, 0)

	for i := 0; i < numCourses; i++ {
		if degrees[i] == 0 {
			coursesAvailable = append(coursesAvailable, i)
		}
	}

	for i := 0; i < len(coursesAvailable); i++ {
		coursesList = append(coursesList, coursesAvailable[i])
		for j := 0; j < len(prereqs[coursesAvailable[i]]); j++ {
			degrees[prereqs[coursesAvailable[i]][j]]--
		}
		for j := 0; j < len(prereqs[coursesAvailable[i]]); j++ {
			if !coursesTaken[prereqs[coursesAvailable[i]][j]] && degrees[prereqs[coursesAvailable[i]][j]] == 0 {
				coursesTaken[prereqs[coursesAvailable[i]][j]] = true
				coursesAvailable = append(coursesAvailable, prereqs[coursesAvailable[i]][j])
			}
		}
	}

	if len(coursesList) < numCourses {
		return make([]int, 0)
	} else {
		return coursesList
	}
}

