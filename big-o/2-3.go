package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//someSlice := []int{5, 4, 0, 2, 1}
	someRand := rand.New(rand.NewSource(time.Now().Unix()))
	someSlice := someRand.Perm(100000)
	minN2(someSlice)
	minN(someSlice)
}

func minN2(arr []int) int {
	defer timeTrack(time.Now(), "N^2")
	min := arr[0]
	smallest := true
	for _, e1 := range arr {
		smallest = true
		for _, e2 := range arr {
			if e1 > e2 {
				smallest = false
			}
		}
		if smallest {
			min = e1
		}
	}

	return min
}

func minN(arr []int) int {
	defer timeTrack(time.Now(), "N")
	min := arr[0]
	for _, e1 := range arr {
		if e1 < min {
			min = e1
		}
	}

	return min
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("%s took %s", name, elapsed)
}
