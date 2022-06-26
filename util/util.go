package util

import (
	"fmt"

	"github.com/JamesHovious/w32"
	"github.com/go-vgo/robotgo"
	"github.com/hnakamur/w32syscall"

	"github.com/samber/lo"
	"log"
	"strings"
	"syscall"
	"time"
)

type Region struct {
	x       int
	y       int
	width   int
	height  int
	centerX int
	centerY int
}

func (r *Region) Center() (int, int) {
	r.centerX = r.x + r.width/2
	r.centerY = r.y + r.height/2
	return r.centerY, r.centerY
}

var (
	LEFT             = "left"
	RIGHT            = "right"
	UP               = "up"
	DOWN             = "down"
	CancelOrAccel    = "esc"
	Enter            = "enter"
	RegionPSWindow   = Region{0, 0, 1788, 1109, 0, 0}
	RegionScreenShot = Region{98, 120, 1587, 891, 0, 0}
	PID              = 0 // ps remote play window pid
)

func L(s string) {
	fmt.Sprintln(" >>> %s", s)
}
func Sleep(second float32) {
	longTime := time.Duration(second) * time.Second
	time.Sleep(longTime)
}
func KeyDown(keyName string) {
	robotgo.KeyDown(keyName)
}
func KeyUp(keyName string) {
	robotgo.KeyUp(keyName)
}
func Press(keyName string) {
	MouseFocus(false)
	robotgo.KeyPress(keyName)
	Sleep(0.2)
}
func BtnCancel() {
	Press(CancelOrAccel)
}
func BtnConfirm() {
	Press(Enter)
}
func Left() {
	Press(LEFT)
}
func Right() {
	Press(RIGHT)
}
func Up() {
	Press(UP)
}
func Down() {
	Press(DOWN)
}

type PosColor struct {
	x        int
	y        int
	hexColor string
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// retry when panic
func AllPixelMatch(posColors []PosColor, fail_call_back_func func()) bool {
	defer duration(track("AllPixelMatch"))
	if len(posColors) <= 0 {
		L("not get  PIXEL pos_and_pixels")
		return false
	}

	isPositionColorEqual := lo.EveryBy[PosColor](posColors, func(pc PosColor) bool {
		newColor := robotgo.GetPixelColor(pc.x, pc.y)
		return newColor == pc.hexColor
	})
	fmt.Println("detect PIXEL found res :", isPositionColorEqual)
	if isPositionColorEqual == false {
		if fail_call_back_func != nil {
			fail_call_back_func()
		}
		AllPixelMatch(posColors, fail_call_back_func)
	} else {
		return true
	}
	return false
}

func MouseFocus(needLog bool) {
	bounds := GetBounds()
	robotgo.Move(bounds.centerX, bounds.centerY)
	robotgo.ActivePID(int32(PID))
	if needLog == true {
		s := fmt.Sprintf("ps remote play window center : %s,%s", bounds.x, bounds.y)
		L(s)
	}
}

func GetBounds() Region {
	x, y, w, h := robotgo.GetBounds(int32(PID))
	fmt.Println("GetBounds is: ", x, y, w, h)
	currentBounds := Region{x, y, w, h, 0, 0}
	currentBounds.Center()
	return currentBounds
}
func FindThanResize() {
	name := "PS Remote Play"
	// find the process id by the process name
	fpid, err := robotgo.FindIds("RemotePlay")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(fpid) <= 0 {
		robotgo.ShowAlert("Error", "NotFound "+name, "Ok", "Cancel")
		return
	}

	pid := fpid[0]
	PID = int(pid)
	fmt.Println(name+" Pid is ", pid)

	err = robotgo.ActivePID(pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	tl := robotgo.GetTitle(pid)
	fmt.Println("title is: ", tl)

	x, y, w, h := robotgo.GetBounds(pid)
	fmt.Println("GetBounds is: ", x, y, w, h)

	err = w32syscall.EnumWindows(func(hwnd syscall.Handle, lparam uintptr) bool {
		h := w32.HWND(hwnd)
		text := w32.GetWindowText(h)
		if strings.Contains(text, name) {
			w32.MoveWindow(h, RegionPSWindow.x, RegionPSWindow.y, RegionPSWindow.width, RegionPSWindow.height, true)
		}
		return true
	}, 0)
	if err != nil {
		log.Fatalln(err)
	}
	//mouse_focus()
	//sleep(1)

	//robotgo.Kill(pid)
}
