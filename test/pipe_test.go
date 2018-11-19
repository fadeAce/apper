package test

import (
	"testing"
	"fmt"
)

func Test_matchAlgorithm(t *testing.T) {
	// 60% of current
	var init int
	init = 200
	for {
		cur := init * 3 / 5
		fmt.Println(cur)
		init = init - cur
		if cur == 1 {
			return
		}
	}
}

func Test_math(t *testing.T) {
	var init int
	init= 2
	fmt.Println(init*3/5)
}
