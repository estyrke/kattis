package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var adj [][]int
var depths []int
var parents []int
var candidates [][]int

func HighestNode(x int) int {
	if parents[x] != x {
		parents[x] = HighestNode(parents[x])
	}
	return parents[x]
}

func JoinLcaSets(parent, child int) {
	parentRoot := HighestNode(parent)
	childRoot := HighestNode(child)

	parents[childRoot] = parentRoot
}

var totalDist = 0

func Dfs(x, src, depth int) {
	depths[x] = depth
	parents[x] = x

	// Check for completed LCAs
	for _, y := range candidates[x] {
		if depths[y] > 0 {
			// y has been visited
			lca := HighestNode(y)
			totalDist += Depth(x) + Depth(y) - 2*Depth(lca) + 1
		}
	}

	// Recurse over children
	for _, node := range adj[x] {
		if node != src {
			Dfs(node, x, depth+1)

			JoinLcaSets(x, node)
		}
	}
}

func Depth(v int) int {
	return depths[v]
}

func main() {
	fmt.Scanln(&n)

	adj = make([][]int, n+1)
	depths = make([]int, n+1)
	parents = make([]int, n+1)
	candidates = make([][]int, n+1)

	ReadEdges()

	BuildCandidates()

	Dfs(1, 0, 1)

	fmt.Println(totalDist)
}

func BuildCandidates() {
	for v1 := 1; v1 <= n/2; v1++ {
		for v2 := v1 * 2; v2 <= n; v2 += v1 {
			candidates[v1] = append(candidates[v1], v2)
			candidates[v2] = append(candidates[v2], v1)
		}
	}
}

func ReadEdges() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n-1; i++ {
		var v1, v2 int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &v1, &v2)
		adj[v1] = append(adj[v1], v2)
		adj[v2] = append(adj[v2], v1)
	}
}
