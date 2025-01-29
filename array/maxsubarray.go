package main

import "fmt"

func MaxSubArray() {
	arsl := []int{0, 0, 1, 0, 1, 1, 1, 0, 0,0,1, 1}
	lastFoundCounter := 0
	mainCounter := 0


	for i := 0; i < len(arsl); i++ {
		if arsl[i] == 1 {
			lastFoundCounter += 1
			if mainCounter < lastFoundCounter {
				mainCounter = lastFoundCounter
			}
		} else {
			lastFoundCounter = 0
		}
	}

	fmt.Println(mainCounter,lastFoundCounter)
}
