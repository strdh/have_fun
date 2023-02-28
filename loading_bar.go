package main

import (
	"fmt"
	"time"
    "strings"
)

func main() {
	for i := 1; i <= 100; i++ {
		// Update progress bar every 100 iterations
		if i%100 == 0 {
			updateProgressBar(i, 100)
		}
		// Do some work here...
		time.Sleep(time.Millisecond * 10)
	}
	// Complete progress bar
	updateProgressBar(100, 100)
}

func updateProgressBar(current, total int) {
	// Calculate percentage complete
	percentage := int((float64(current) / float64(total)) * 100)
	// Calculate number of completed blocks
	completedBlocks := int(float64(percentage) / 2)
	// Output progress bar
	fmt.Printf("\r[%s%s] %d%%", strings.Repeat("=", completedBlocks), strings.Repeat(" ", 50-completedBlocks), percentage)
}
