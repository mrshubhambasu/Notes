package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

var vowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

// Optimized and parallelized version with tracing enabled
func countVowelsParallel(s string) int {
	// Create a file to store the trace data
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("Failed to create trace file:", err)
		return -1
	}
	defer f.Close()

	// Start tracing
	err = trace.Start(f)
	if err != nil {
		fmt.Println("Failed to start trace:", err)
		return -1
	}
	defer trace.Stop()

	// Start parallel vowel counting
	var wg sync.WaitGroup
	ch := make(chan int)

	numGoroutines := 4
	partSize := len(s) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := (i + 1) * partSize
		if i == numGoroutines-1 {
			end = len(s)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			count := 0
			for _, ch := range s[start:end] {
				if vowels[ch] {
					count++
				}
			}
			ch <- count
		}(start, end)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	total := 0
	for count := range ch {
		total += count
	}

	return total
}

func main() {
	// Example string for vowel counting
	s := "This is a simple example string to count vowels using goroutines."

	// Set the number of CPUs to use (optional)
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Run the function with tracing enabled
	vowelCount := countVowelsParallel(s)

	fmt.Printf("Vowel count: %d\n", vowelCount)

	// After running, you can analyze the trace with:
	// go tool trace trace.out
}
