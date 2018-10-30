package main

func bubbleSort(valuesArr []int) []int {
	for i := 0; i < len(valuesArr)-1; i++ {
		for j := len(valuesArr) - 1; j > i; j-- {
			if valuesArr[j] < valuesArr[j-1] {
				temp := valuesArr[j]
				valuesArr[j] = valuesArr[j-1]
				valuesArr[j-1] = temp
			}
		}
	}
	return valuesArr
}
