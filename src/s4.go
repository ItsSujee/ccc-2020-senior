package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func countDiff(a string, b string) int {
	output := 0
	for i:=0; i<len(a); i++{
		if a[i] != b[i] {
			output++
		}
	}
	return output
}

func main() {
	var input string
	fmt.Scanf("%s", &input)

	var windows []string
	for i:=0; i<len(input); i++{
		buffer := input[i:] + input[0:i]
		windows = append(windows, buffer)
	}

	inputList := strings.Split(input, "")
	sort.Strings(inputList)
	sortedInput := strings.Join(inputList, "")

	var diff []int
	for _,v := range windows{
		diff = append(diff, countDiff(v, sortedInput))
	}

	sort.Ints(diff)
	output := math.Ceil(float64(diff[0])/2)
	fmt.Println(output)
}
