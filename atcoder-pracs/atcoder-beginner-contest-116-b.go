package main

import (
	"fmt"
	"os"
)

func f(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	s := 0
	fmt.Fscanf(os.Stdin, "%d", &s)

	a := make([]int, 1000000)
	a[0] = s
	for i := 1; i < 1000000; i++ {
		a[i] = f(a[i-1])
	}
	for i := 0; i < 1000000; i++ {
		for j := i + 1; j < 1000000; j++ {
			if a[i] == a[j] {
				fmt.Fprintf(os.Stdout, "%d\n", j+1)
				return
			}
		}
	}
}
