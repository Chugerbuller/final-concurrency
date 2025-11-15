package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	rand.Seed(time.Now().UnixNano())

	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(10000) // диапазон можно менять
	}

	return result
}

// maximum returns the maximum number of elements.
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

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	chunkSize := SIZE / CHUNKS
	var max atomic.Int32
	max.Store(math.MinInt32)
	maximums := make([]int, 8)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		go func(part []int) {
			defer wg.Done()
			localMax := maximum(data)
			maximums = append(maximums, localMax)
		}(data[start:end])
	}
	wg.Wait()
	return maximum(maximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	end := time.Now()
	elapsed := time.Duration(end.Sub(start).Milliseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	end = time.Now()
	elapsed = time.Duration(end.Sub(start).Milliseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
