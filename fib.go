package main

import (
	"fmt"
	"os"
	"strconv"
)

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("usage: fib [integer]\n")
		os.Exit(1)
	}

	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "argument must be integer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d", Fib(i))
}