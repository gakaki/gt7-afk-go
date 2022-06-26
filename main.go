package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.gakaki.com/gt7-afk-go/game"
	u "github.gakaki.com/gt7-afk-go/util"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println(" from Windows")
		u.FindThanResize()
	}

	// find the process id by the process name
	fpid, err := robotgo.FindIds("RemotePlay")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fpid)
	game.Reward117()
}
