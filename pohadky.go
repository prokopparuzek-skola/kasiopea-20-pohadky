package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		var isRead []bool
		var sortedBook [][2]int
		var words int
		fmt.Scan(&N)
		sortedBook = make([][2]int, N)
		isRead = make([]bool, N)
		for j := range isRead {
			isRead[j] = true
		}
		for j := 0; j < N; j++ {
			var b int
			fmt.Scan(&b)
			sortedBook[j] = [2]int{b, j}
			words += b
		}
		sort.Slice(sortedBook, func(i, j int) bool { return sortedBook[i][0] > sortedBook[j][0] })
		for _, s := range sortedBook {
			switch s[1] {
			case 0:
			case N - 1:
				if isRead[N-2] {
					isRead[N-1] = false
					words -= s[0]
				}
			default:
				if isRead[s[1]-1] && isRead[s[1]+1] {
					isRead[s[1]] = false
					words -= s[0]
				}
			}
		}
		fmt.Println(words)
	}
}
