package dateRoutine

import "math"

// GetDistance returns distance in days between two dates
func GetDistance(d1, d2, m1, m2, y1, y2 int) int {
	return int(math.Abs(float64((y1-y2)*365 + (m1-m2)*30 + (d1 - d2))))
}
