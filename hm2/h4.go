package main

import "fmt"

func main() {
	const N = 10000
	var arr [N - 1]int
	for i := 2; i <= N; i++ {
		arr[i-2] = i
	}

	for i := 0; i < N-1; i++ {
		var p int = 0
		if arr[i] != 0 {
			p = arr[i]
		} else {
			continue
		}

		for j := i + 1; j < N-1; j++ {
			if arr[j]%p == 0 {
				arr[j] = 0
			}
		}
	}

	var resultArr [100]int
	var resultArrCount int
	for i := 0; i < N-1; i++ {
		if arr[i] != 0 {
			resultArr[resultArrCount] = arr[i]
			resultArrCount++
			if resultArrCount == 100 {
				break
			}
		}
	}
	fmt.Println(resultArr)
}
