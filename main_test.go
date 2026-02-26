package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("size zero", func(t *testing.T) {
		arr := generateRandomElements(0)
		assert.Empty(t, arr)
	})

	t.Run("size negative", func(t *testing.T) {
		arr := generateRandomElements(-10)
		assert.Empty(t, arr)
	})

	t.Run("size positive", func(t *testing.T) {
		size := 1000
		arr := generateRandomElements(size)
		assert.Len(t, arr, size)
	})
}

// Тесты maxChunks с использованием generateRandomElements
func TestMaxChunks(t *testing.T) {
	t.Run("large random array 1M", func(t *testing.T) {
		data := generateRandomElements(1000000)
		expectedMax := maximum(data)
		assert.Equal(t, expectedMax, maxChunks(data))
	})

	t.Run("large random array 500k", func(t *testing.T) {
		data := generateRandomElements(500000)
		expectedMax := maximum(data)
		assert.Equal(t, expectedMax, maxChunks(data))
	})

	t.Run("small random array 100", func(t *testing.T) {
		data := generateRandomElements(100)
		expectedMax := maximum(data)
		assert.Equal(t, expectedMax, maxChunks(data))
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		data := []int{}
		assert.Equal(t, 0, maximum(data))
	})
	t.Run("one element", func(t *testing.T) {
		data := []int{42}
		assert.Equal(t, 42, maximum(data))
	})

	t.Run("all negative", func(t *testing.T) {
		data := []int{-10, -3, -50, -1}
		assert.Equal(t, -1, maximum(data))
	})
	t.Run("all same", func(t *testing.T) {
		data := []int{7, 7, 7, 7, 7}
		assert.Equal(t, 7, maximum(data))
	})
	t.Run("normal array", func(t *testing.T) {
		data := []int{1, 5, 3, 9, 2, 8}
		assert.Equal(t, 9, maximum(data))
	})
}
