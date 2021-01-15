package main

import "fmt"

func CardScore(card string, trump byte) int {
	switch {
	case card[0] == 'A':
		return 11
	case card[0] == 'K':
		return 4
	case card[0] == 'Q':
		return 3
	case card[0] == 'J' && card[1] == trump:
		return 20
	case card[0] == 'J':
		return 2
	case card[0] == 'T':
		return 10
	case card[0] == '9' && card[1] == trump:
		return 14
	default:
		return 0
	}

}

func main() {
	var n int
	var b string

	nread, err := fmt.Scanln(&n, &b)
	if nread != 2 {
		fmt.Println(err)
	}

	score := 0

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			var card string
			fmt.Scan(&card)
			score += CardScore(card, b[0])
		}
	}
	fmt.Println(score)
}
