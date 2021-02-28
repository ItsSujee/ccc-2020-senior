package main

import "fmt"

func heapPermutation(a []int32, size int, n int, ptr *[]string) {
	if size == 1 {
		*ptr = append(*ptr, string(a))
		return
	}
	for i:=0; i<size; i++ {
		heapPermutation(a, size-1, n, *&ptr)
		if size % 2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func dropDuplicate(a []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	var N string
	var H string
	fmt.Scanf("%s", &N)
	fmt.Scanf("%s", &H)

	var permutation []string
	heapPermutation([]int32(N), len(N), len(N), &permutation)
	permutation = dropDuplicate(permutation)

	numWindows := len(H) - len(N)
	var windows []string
	for i:=0; i<=numWindows; i++{
		windows = append(windows, H[i:i+len(N)])
	}
	windows = dropDuplicate(windows)

	var counter = 0
	for _,v1 := range windows{
		for _,v2 := range permutation{
			if v1 == v2 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
