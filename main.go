package main

import (
	"fmt"
	"github.gakaki.com/gt7-afk-go/game"
	u "github.gakaki.com/gt7-afk-go/util"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println(" from Windows")
		u.FindThanResize()
		game.Reward117()
	}

}
