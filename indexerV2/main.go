package main

import (
	"fmt"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/cmd"
)

func main() {

	startTime := time.Now()

	cmd.Start()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := duration.Seconds()
	minutes := duration.Minutes()
	fmt.Printf("The code ran in %.2f seconds\n", seconds)
	fmt.Printf("The code ran in %.2f minutes\n", minutes)
}
