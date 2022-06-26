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
//PosColor {506, 330, "0b0a0b"}
//PosColor {511, 340, "ffffff"}
//PosColor {526, 329, "0a0a0a"}
