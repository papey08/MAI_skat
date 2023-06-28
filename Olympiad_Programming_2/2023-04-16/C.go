package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	pr := sieve(n)
	ans := 0
	for i := 2; i < n; i++ {
		if pr[i] {
			for t := i; t+i < n; t++ {
				if n-i-2*t >= 0 && pr[n-i-t] && pr[t] {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

func sieve(n int) []bool {
	res := make([]bool, n)
	for i := 2; i < len(res); i++ {
		res[i] = true
	}

	for i := 2; i*i < len(res); i++ {
		if res[i] {
			for j := i * i; j < len(res); j += i {
				res[j] = false
			}
		}
	}
	return res
}
