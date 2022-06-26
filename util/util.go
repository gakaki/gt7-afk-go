package util

import (
	"fmt"
	"github.com/JamesHovious/w32"

	//"github.com/JamesHovious/w32"

	//"github.com/JamesHovious/w32"
	"runtime"

	"github.com/go-vgo/robotgo"
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
	PID              = 0 // ps remote play window pid\
)

//var PSHWND w32.HWND

func Sleep(t time.Duration) {
	time.Sleep(t)
}
func KeyDown(keyName string) {
	robotgo.KeyDown(keyName)
}
func KeyUp(keyName string) {
	robotgo.KeyUp(keyName)
}
func Press(keyName string) {
	MouseFocus(false)
	fmt.Println("key press:", keyName)
	KeyDown(keyName)
	Sleep(200 * time.Millisecond)
	KeyUp(keyName)
	Sleep(700 * time.Millisecond)
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

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

type PosColor struct {
	x        int
	y        int
	hexColor string
}

// retry when panic
func AllPixelMatch(posColors []PosColor, fail_call_back_func func()) bool {
	defer duration(track("AllPixelMatch"))
	if len(posColors) <= 0 {
		fmt.Println("not get  PIXEL pos_and_pixels")
		return false
	}
	boolRes := make([]bool, 0)
	finalRes := true
	for _, row := range posColors {
		newColor := robotgo.GetPixelColor(row.x, row.y)
		res := (newColor == row.hexColor)
		boolRes = append(boolRes, res)
		if res == false {
			finalRes = false
		}
	}

	fmt.Println("detect PIXEL found res :", finalRes, boolRes)
	if finalRes == false {
		if fail_call_back_func != nil {
			fail_call_back_func()
		}
		AllPixelMatch(posColors, fail_call_back_func)
	} else {
		return true
	}
	return false
}

func WaitCafeMenuPixel() bool {
	pos_and_pixels := []PosColor{
		PosColor{1108, 686, "ffffff"},
		//PosColor{1130, 683, "ffffff"},
	}
	if AllPixelMatch(pos_and_pixels, func() {
		fmt.Println("detect PIXEL_CAFE found res : False")
		Sleep(500 * time.Millisecond)
	}) == true {
		fmt.Println(("detect PIXEL_CAFE found res : True"))
		Sleep(500 * time.Millisecond)
		return true
	} else {
		return false
	}
}

func WaitCarHomePixel() bool {
	pos_and_pixels := []PosColor{
		PosColor{1479, 699, "ffffff"},
	}
	if AllPixelMatch(pos_and_pixels, func() {
		fmt.Println("detect PIXEL_CAR_HOME found res : False")
		Sleep(500 * time.Millisecond)
	}) == true {
		fmt.Println(("detect PIXEL_CAR_HOME found res : True"))
		Sleep(500 * time.Millisecond)
		return true
	} else {
		return false
	}
}

func WaitAlreadyGetRewardPixel() bool {
	pos_and_pixels := []PosColor{
		PosColor{511, 484, "ffffff"},
	}

	if AllPixelMatch(pos_and_pixels, func() {
		fmt.Println("detect IMG_ALREADY_GOT found res : False")
		BtnConfirm()
		Sleep(500 * time.Millisecond)
	}) == true {
		fmt.Println(("detect IMG_ALREADY_GOT found res : True"))
		Sleep(500 * time.Millisecond)
		return true
	} else {
		return false
	}
}

func MouseFocus(needLog bool) {
	bounds := GetBounds()
	robotgo.Move(bounds.centerX, bounds.centerY)

	if runtime.GOOS == "windows" {
		w32.SetFocus(PSHWND)
		w32.SetForegroundWindow(PSHWND)
		robotgo.SetFocus(win.HWND(PSHWND))
	}

	robotgo.Click()
	if needLog == true {
		fmt.Println("ps remote play window center : ", bounds.x, bounds.y)
	}
}

func GetBounds() Region {

	rect := w32.GetWindowRect(PSHWND)
	width := rect.Right - rect.Left
	height := rect.Bottom - rect.Top

	currentBounds := Region{int(rect.Left), int(rect.Top), int(width), int(height), 0, 0}
	currentBounds.Center()
	//fmt.Println("Bounds is: ", currentBounds)
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

	// no effect for ps remote play
	//x, y, w, h := robotgo.GetBounds(pid)
	//fmt.Println("GetBounds is: ", x, y, w, h)

	if runtime.GOOS == "windows" {
		err = w32syscall.EnumWindows(func(hwnd syscall.Handle, lparam uintptr) bool {
			h := w32.HWND(hwnd)

			text := w32.GetWindowText(h)
			if strings.Contains(text, name) {
				PSHWND = h
				w32.MoveWindow(h, RegionPSWindow.x, RegionPSWindow.y, RegionPSWindow.width, RegionPSWindow.height, true)
				GetBounds()
			}
			return true
		}, 0)
		if err != nil {
			log.Fatalln(err)
		}
		MouseFocus(false)
	}

	//sleep(1)
	//robotgo.Kill(pid)
}
