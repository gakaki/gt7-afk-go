package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func addMouse() {
	fmt.Println("--- Start Capture Mouse Pos RGB ---")
	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		color := robotgo.GetPixelColor(int(e.X), int(e.Y))
		s := fmt.Sprintf("PosColor {%v, %v, \"%v\"}", e.X, e.Y, color)
		fmt.Println(s)
	})

	defer hook.End()
	s := hook.Start()

	<-hook.Process(s)
}

func main() {
	addMouse()
}

//cafe menu
//PosColor {888, 553, "ffffff"}
//PosColor {875, 524, "151515"}

//home
//PosColor {1172, 556, "ffffff"}
//PosColor {1158, 543, "151515"}
