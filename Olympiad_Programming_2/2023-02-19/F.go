package main

import (
	"fmt"
	"strconv"
	"sync"
)

func check(n string, prime string) (bool, []int) {
	var p1, p2 int
	var ans []int
	isPossible := false
	for p1 < len(n) {
		if n[p1] == prime[p2] {
			if p2 == len(prime)-1 {
				isPossible = true
				p1++
				for i := p1; i < len(n); i++ {
					ans = append(ans, i)
				}
				break
			} else {
				p1++
				p2++
				continue
			}
		} else {
			ans = append(ans, p1)
			p1++
		}
	}
	return isPossible, ans
}

func main() {
	primes := []string{"2", "3", "5", "7", "11", "19", "41", "61", "89", "409", "449", "499", "881", "991", "6469", "6949", "9001", "9049", "9649", "9949", "60649", "666649", "946669", "60000049", "66000049", "66600049"}
	var maxPrime int
	var maxAns []int
	var n string
	fmt.Scan(&n)
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	for _, p := range primes {
		wg.Add(1)
		go func(prime string) {
			defer wg.Done()
			ok, ans := check(n, prime)
			if ok {
				temp, _ := strconv.Atoi(prime)
				mu.Lock()
				if maxPrime < temp {
					maxPrime = temp
					maxAns = ans
				}
				mu.Unlock()
			}
		}(p)
	}
	wg.Wait()
	fmt.Println(maxPrime)
	if len(maxAns) == 0 {
		fmt.Println("-1")
	} else {
		for _, x := range maxAns {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}
