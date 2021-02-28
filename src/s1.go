package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var i int
	var N int
	fmt.Scanf("%d", &N)

	var T = make([]int, N)
	var X = make([]int, N)
	for i:=0;i<N;i++{
		fmt.Scanf("%d", &T[i])
		fmt.Scanf("%d", &X[i])
	}

	//Sort T
	var sortedT = make([]int, N)
	for index,value := range T {
		sortedT[index] = value
	}
	sort.Ints(sortedT)

	//X given sorted T
	var sortedX = make([]int, N)
	for i=0; i<N; i++{
		for index, value := range T {
			if sortedT[i] == value {
				sortedX[i] = X[index]
			}
		}
	}

	//Speed Calculation
	var speed = make([]float64, N-1)
	for i:=0;i<N-1;i++{
		buffer := float64(sortedX[i+1] - sortedX[i]) / float64(sortedT[i+1] - sortedT[i])
		speed[i] = math.Abs(buffer)
	}

	//Max speed
	var output float64
	for i:=0; i<N-1; i++{
		if speed[i] > output{
			output = speed[i]
		}
	}

	fmt.Println(output)
}
