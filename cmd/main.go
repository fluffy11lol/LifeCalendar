package main

import (
	dateRoutine "LifeCounter/pkg/dateRoutine"
	printer "LifeCounter/pkg/printing"
	user "LifeCounter/pkg/userInitializing"
	"fmt"
	"time"
)

func main() {
	var curUser user.User
	err := curUser.InitInfoFromJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !dateRoutine.ValidateBirthDate(curUser.Birthdate.Day, curUser.Birthdate.Month, curUser.Birthdate.Year) {
		fmt.Printf("Error: Incorrect birthdate")
		return
	}
	curDate := time.Now()
	curDistance := dateRoutine.GetDistance(curUser.Birthdate.Day, curDate.Day(), curUser.Birthdate.Month,
		int(curDate.Month()), curUser.Birthdate.Year, curDate.Year())
	printer.PrintLifeCalendar(curDistance)
	printer.PrintGreeting(curUser.Name)
}
