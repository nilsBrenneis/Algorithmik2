package main

import (
	"fmt"
	"math"
)

func isPrime(value int) bool {
	n := float64(value)
	prim := true
	k := 2.0
	for k < n {
		if math.Mod(n, k) == 0 {
			prim = false
		}
		k++
	}
	return prim
}

func binarySearch(valuesArr []int, key int) int {
	low := 0
	high := len(valuesArr)
	for low <= high {
		mid := int((low + high) / 2)
		if valuesArr[mid] < key {
			low = mid + 1
		} else if valuesArr[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func bubblesort(valuesArr []int) []int {
	for i := len(valuesArr); i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if valuesArr[j] > valuesArr[j+1] {
				temp := valuesArr[j+1]
				valuesArr[j+1] = valuesArr[j]
				valuesArr[j] = temp
			}
		}
	}
	return valuesArr
}

func main() {
	testData := []int{1, 5, 89, 56, 44, 2, 0}
	fmt.Println("Bubble Sort:", bubblesort(testData))
	fmt.Println("Binary Search 56:", binarySearch(testData, 56))

	prime := 408
	fmt.Println(prime, "isPrime?:", isPrime(prime))
	prime = 409
	fmt.Println(prime, "isPrime?:", isPrime(prime))
}
