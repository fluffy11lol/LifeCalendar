package printing

import "fmt"

const passLife = "🟩"
const futureLife = "🟦"

func PrintLifeCalendar(curDistance int) {
	maxMonth := 69*365/30 + 1
	curMonth := curDistance / 30
	for i := 1; i <= maxMonth; i++ {
		if i < curMonth {
			fmt.Print(passLife)
		} else {
			fmt.Print(futureLife)
		}
		if i%40 == 0 {
			fmt.Println()
		}
	}
}
