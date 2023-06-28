package main

import "fmt"

const (
	N   = 501
	mod = 1000000007
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	arr := make([][]int, N)
	for i := range arr {
		arr[i] = make([]int, N)
	}
	arr[0][0] = 1

	for i := range arr {
		for j := range arr {
			if i >= 3 {
				arr[i][j] += arr[i-3][j]
			}
			if i >= 2 {
				arr[i][j] += arr[i-2][j]
			}
			if j >= 3 {
				arr[i][j] += arr[i][j-3]
			}
			if j >= 2 {
				arr[i][j] += arr[i][j-2]
			}
			arr[i][j] %= mod
		}
	}
	fmt.Println(arr[a][b])
}
