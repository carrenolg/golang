package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	// Example array of integers
	arr := []int{38, 27, 43, 3, 9, 82, 10, 15, 22, 17, 56, 72, 91, 44, 33, 66}

	// Partition the array into 4 parts
	partitions := partitionArray(arr, 4)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(len(partitions))

	// Sort each partition concurrently
	for i := range partitions {
		go func(i int) {
			defer wg.Done()
			sort.Ints(partitions[i])
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Merge the sorted partitions
	sortedArr := mergeSortedPartitions(partitions)

	// Print the final sorted array
	fmt.Println("Sorted Array:", sortedArr)
}

// partitionArray divides the array into n parts of approximately equal size
func partitionArray(arr []int, n int) [][]int {
	partitions := make([][]int, n)
	size := len(arr) / n
	remainder := len(arr) % n

	start := 0
	for i := 0; i < n; i++ {
		end := start + size
		if i < remainder {
			end++
		}
		partitions[i] = arr[start:end]
		start = end
	}

	return partitions
}

// mergeSortedPartitions merges the sorted partitions into one sorted array
func mergeSortedPartitions(partitions [][]int) []int {
	// Initialize pointers for each partition
	pointers := make([]int, len(partitions))
	result := []int{}

	for {
		// Find the smallest element among the current elements of each partition
		minVal := -1
		minIndex := -1
		for i := range partitions {
			if pointers[i] < len(partitions[i]) {
				if minIndex == -1 || partitions[i][pointers[i]] < minVal {
					minVal = partitions[i][pointers[i]]
					minIndex = i
				}
			}
		}

		// If no more elements to merge, break the loop
		if minIndex == -1 {
			break
		}

		// Append the smallest element to the result and move the pointer
		result = append(result, minVal)
		pointers[minIndex]++
	}

	return result
}
