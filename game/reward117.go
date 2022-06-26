package game

import (
	"fmt"
	u "github.gakaki.com/gt7-afk-go/util"
	"time"
)

func Reward117() {
	start := time.Now()

	Reward(false)
	ToGarage()

	Reward(true)
	ToGarage()

	elapsed := time.Since(start)
	fmt.Sprintln(" reward117 usage  %s", elapsed)

	Reward117()
}

func ToGarage() {
	u.WaitCafeMenuPixel()

	u.Right()
	u.BtnConfirm()
	//loading into  car home
	u.Sleep(7*time.Second + 500*time.Millisecond)

	u.Right()
	u.Right()
	u.Right()

	//into get reward page
	u.BtnConfirm()
	u.Sleep(500 * time.Millisecond)
	u.BtnConfirm()
	u.BtnConfirm()

	// wait the reward animation
	u.Sleep(11 * time.Second)
	u.WaitAlreadyGetRewardPixel()

	u.BtnCancel()
	u.BtnCancel()

	u.Sleep(4 * time.Second)
	u.WaitCarHomePixel()

	u.Left()
}
func Reward(isEngine bool) {

	u.BtnConfirm()

	u.Sleep(3*time.Second + 500*time.Millisecond)

	//奖杯
	u.Left()

	u.BtnConfirm()

	u.Down()

	u.Right()

	u.BtnConfirm()

	//append menu page
	u.Up()

	// rotate engine
	if isEngine == true {
		u.Right()
		u.Right()
	} else { //toyota 86

	}

	u.BtnConfirm()
	u.Sleep(500 * time.Millisecond)
	u.BtnConfirm()
	u.Sleep(700 * time.Millisecond)

	u.BtnCancel()
	u.BtnCancel()
	u.BtnCancel()

	//loading time for return to dashboard
	u.Sleep(5 * time.Second)
}
