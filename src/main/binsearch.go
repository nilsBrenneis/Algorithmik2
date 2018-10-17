package main

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
