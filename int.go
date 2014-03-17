package scanutil

import (
	"fmt"
)

func IntLine(n int) []int {
	p := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scanf("%d", &p[j])
	}
	fmt.Scanln()
	return p
}