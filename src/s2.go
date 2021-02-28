package main

import "fmt"

func findValue(grid [][]int, value int) [][]int {
	var i,j int
	var output [][]int
	M := len(grid)
	N := len((grid)[0])
	for i=0; i<M; i++{
		for j=0;j<N; j++{
			if (grid)[i][j] == value {
				buffer := []int {i,j}
				output = append(output, buffer)
			}
		}
	}
	return output
}

func recursiveSearch(grid [][] int, product [][]int, search int) string {
	productIndex := findValue(product, search)
	var newSearch []int
	for i:=0; i<len(productIndex);i++{
		buffer := grid[productIndex[i][0]][productIndex[i][1]]
		newSearch = append(newSearch, buffer)
	}
	for j:=0; j<len(newSearch);j++{
		if newSearch[j] == -1 {
			return "yes"
		}
		return recursiveSearch(grid, product, newSearch[j])
	}
	return "no"
}


func main() {
	var buffer int
	var i,j int

	var M int
	var N int
	fmt.Scanf("%d\n",&M)
	fmt.Scanf("%d\n",&N)

	//Input M x N integers
	var grid = make([][]int, M)
	for i:=0; i<M; i++{
		grid[i] = make([]int, N)
	}
	for i=0; i<M; i++{
		for j=0; j<N; j++{
			fmt.Scanf("%d", &buffer)
			grid[i][j] = buffer
		}
	}

	//Set target
	grid[M-1][N-1] = -1

	//Create M,N product matrix
	var product = make([][]int, M)
	for i:=0; i<M; i++{
		product[i] = make([]int, N)
	}
	for i=1; i<M+1; i++{
		for j=1; j<N+1; j++{
			product[i-1][j-1] = i * j
		}
	}

	start := grid[0][0]
	ans := recursiveSearch(grid, product, start)
	fmt.Println(ans)
}