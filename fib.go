package main

import (
	"fmt"
	"github.com/codegangsta/cli"
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
	app := cli.NewApp()
	app.Name = "fibonacci"
	app.Usage = "print arguments"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 1 {
			fmt.Print("usage: fib [integer]\n")
			os.Exit(1)
		}

		i, err := strconv.Atoi(c.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "argument must be integer: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("%d", Fib(i))
	}

	app.Run(os.Args)
}