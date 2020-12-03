package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		var book []int
		var words int
		fmt.Scan(&N)
		book = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&book[j])
		}
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
				words += book[j]
				if book[j+1] < book[j+2] {
					words += book[j+1]
					break loop
				}
				j += 2
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
