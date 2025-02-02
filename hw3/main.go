package main

import "fmt"

func main() {

	PrintChessBoard(8, 8)
}

func PrintChessBoard(a int, b int) {
	var s1 string
	for i := 1; i <= a; i++ {

		for j := 1; j <= b; j++ {
			if !(i%2 == 0) {
				if j%2 == 0 {
					s1 = " "
				} else {
					s1 = "#"
				}
			} else {
				if j%2 == 0 {
					s1 = "#"
				} else {
					s1 = " "
				}
			}

			fmt.Print(s1)

		}
		fmt.Println()
	}
}
