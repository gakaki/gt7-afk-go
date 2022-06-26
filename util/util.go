package util

import (
	"fmt"

	"github.com/JamesHovious/w32"
	"github.com/hnakamur/w32syscall"

	"github.com/go-vgo/robotgo"
	"log"
	"strings"
	"syscall"
	"time"
)

type Region struct {
	x      int
	y      int
	width  int
	height int
}

var (
	LEFT             = "left"
	RIGHT            = "right"
	UP               = "up"
	DOWN             = "down"
	CancelOrAccel    = "esc"
	Enter            = "enter"
	RegionPSWindow   = Region{0, 0, 1788, 1109}
	RegionScreenShot = Region{98, 120, 1587, 891}
)

func Sleep(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}
func KeyDown(keyName string) {
	robotgo.KeyDown(keyName)
}
func KeyUp(keyName string) {
	robotgo.KeyUp(keyName)
}
func MouseFocus(needLog bool) {

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
