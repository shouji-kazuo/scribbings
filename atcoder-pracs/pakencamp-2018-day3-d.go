package main

import (
	"fmt"
	"os"
)

func main() {
	N := 0
	fmt.Fscanf(os.Stdin, "%d", &N)

	C := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(os.Stdin, "%d", &(C[i]))
	}

	X := make([]int, N-1)
}
