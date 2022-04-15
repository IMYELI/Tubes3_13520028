package Script

import "math"

func LCSDistance(closedMatch string, DNApenyakit string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if closedMatch[m-1] == DNApenyakit[n-1] {
		return 1 + LCSDistance(closedMatch, DNApenyakit, m-1, n-1)
	} else {
		return int(math.Max(float64(LCSDistance(closedMatch, DNApenyakit, m, n-1)),
			float64(LCSDistance(closedMatch, DNApenyakit, m-1, n))))
	}
}

func Similarity(length int, distance int) int {
	return distance * 100 / length
}
