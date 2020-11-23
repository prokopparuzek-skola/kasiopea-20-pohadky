package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		var book []int
		fmt.Scan(&N)
		book = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&book[j])
		}
		var words int
	loop:
		for j := 0; ; {
			switch N - j {
			case 0:
				break loop
			case 1:
				words += book[j]
				break loop
			case 2:
				words += book[j]
				break loop
			case 3:
				if book[j+1] < book[j+2] {
					words += book[j] + book[j+1]
				} else {
					words += book[j] + book[j+2]
				}
				break loop
			default:
				words += book[j]
				if book[j+1]+book[j+3] < book[j+2] {
					words += book[j+1]
					j += 3
				} else {
					j += 2
				}
			}
		}
		fmt.Println(words)
	}
}
