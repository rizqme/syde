package main

import (
	"fmt"
	"os"

	"github.com/feedloop/syde/internal/dashboard"
)

func main() {
	if err := dashboard.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "syded: %v\n", err)
		os.Exit(1)
	}
}
