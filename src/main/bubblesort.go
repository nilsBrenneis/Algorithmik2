package main

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
