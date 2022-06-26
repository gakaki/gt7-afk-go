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
		fmt.Sprintln("{%v, %v, \"%v\"}", e.X, e.Y, color)
	})

	defer hook.End()
	s := hook.Start()

	<-hook.Process(s)
}

func main() {
	addMouse()
}
