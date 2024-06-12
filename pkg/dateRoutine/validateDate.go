package dateRoutine

import "time"

func ValidateBirthDate(d, m, y int) bool {
	if d > 0 && d < 31 && m > 0 && m <= 12 && y > 1900 && y <= time.Now().Year() {
		return true
	} else {
		return false
	}
}
