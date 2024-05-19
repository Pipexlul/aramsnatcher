package main

import (
	"fmt"

	"github.com/ImOlli/go-lcu/lcu"
)

func logFmtNewline(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	conInfo, err := lcu.FindLCUConnectInfo()
	if err != nil {
		logFmtNewline("Error: %w", err)
		return
	}
}
