package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// Returns a slice of the given size filled
// with pseudo-random integers. If size is non-positive, it returns
// an empty slice.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	src := rand.NewSource(time.Now().UnixNano())

	result := make([]int, size)
	for i := range size {
		result[i] = int(src.Int63()) // диапазон можно менять
	}

	return result
}

// Returns the largest value in the slice.
// If the slice is empty, it returns 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// Returns the maximum value in the given slice using concurrent processing.
// If the slice is empty, it returns 0.
func maxChunks(data []int) int {
	length := len(data)

	if length == 0 {
		return 0
	}

	chunkSize := (length + CHUNKS - 1) / CHUNKS
	results := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := range CHUNKS {
		start := min(i*chunkSize, length)
		end := min(start+chunkSize, length)

		wg.Add(1)

		go func(chunk []int, idx int) {
			defer wg.Done()
			results[idx] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()

	globalMax := math.MinInt
	for i := range CHUNKS {
		if results[i] > globalMax {
			globalMax = results[i]
		}
	}

	return globalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	end := time.Now()
	diff := end.Sub(start).Milliseconds()
	elapsed := time.Duration(diff)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	end = time.Now()
	diff = end.Sub(start).Milliseconds()
	elapsed = time.Duration(diff)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
