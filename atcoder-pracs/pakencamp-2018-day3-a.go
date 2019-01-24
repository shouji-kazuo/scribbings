package main

import (
	"fmt"
	"os"
)

func main() {
	y, m, d := 0, 0, 0
	fmt.Fscanf(os.Stdin, "%d %d %d", &y, &m, &d)

	base := 2018
	if m == 12 && d == 25 {
		fmt.Fprintf(os.Stdout, "%d\n", y-base)
		return
	}
	fmt.Fprintln(os.Stdout, "NOT CHRISTMAS DAY")
}
