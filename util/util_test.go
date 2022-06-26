package util

import (
	"fmt"
	"testing"
)

func TestWaitAlreadyGetRewardPixel(t *testing.T) {
	FindThanResize()
	fmt.Println(WaitAlreadyGetRewardPixel())
}

func TestWaitCarHomePixel(t *testing.T) {
	FindThanResize()
	fmt.Println(WaitCarHomePixel())
}

func TestWaitCafeMenuPixel(t *testing.T) {
	FindThanResize()
	fmt.Println(WaitCafeMenuPixel())
}
