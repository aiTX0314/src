package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func f(num int) int {
	var min float64 = 1
	var max float64 = 100

	for math.Ceil(min) != math.Floor(max) {
		if (max+min)/2 > float64(num) {
			max = (max + min) / 2
		} else {
			min = (max + min) / 2
		}
	}

	return int(max)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var num int = rand.Intn(100)
	//fmt.Println(num)
	fmt.Println("随机数为:", f(num))
}
