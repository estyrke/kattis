package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FindMaxHeight(heights []int) int {
	var maxHeight = 0
	var leftAnchorHeight = 0
	var rightAnchorHeights = make([]int, len(heights))
	var n = len(heights)

	if n < 3 {
		return 0
	}

	rightAnchorHeights[n-2] = heights[n-1]
	for i := len(heights) - 2; i >= 0; i-- {
		var rightHeight = heights[i+1]
		if rightHeight > rightAnchorHeights[i+1] {
			rightAnchorHeights[i] = rightHeight
		} else {
			rightAnchorHeights[i] = rightAnchorHeights[i+1]
		}
	}

	for i := 0; i < len(heights); i++ {
		var height = heights[i]
		if height > leftAnchorHeight {
			leftAnchorHeight = height
			continue
		}
		var rightAnchorHeight = rightAnchorHeights[i]
		var jumpHeight = leftAnchorHeight - height
		if rightAnchorHeight < leftAnchorHeight {
			jumpHeight = rightAnchorHeight - height
		}

		if jumpHeight > maxHeight {
			maxHeight = jumpHeight
		}
	}
	return maxHeight
}

func main() {
	// f, err := os.Create("bungeebuilder.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	var maxHeight = 0
	var n int

	fmt.Scanln(&n)

	var heights = make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		height, _ := strconv.Atoi(scanner.Text())
		heights[i] = height
	}

	maxHeight = FindMaxHeight(heights)

	fmt.Println(maxHeight)
}
