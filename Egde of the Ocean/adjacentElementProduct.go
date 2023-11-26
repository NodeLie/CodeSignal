package main

func solution(inputArray []int) int {
	maxSolution := inputArray[0] * inputArray[1]
	for i := 0; i < len(inputArray)-1; i++ {
		product := inputArray[i] * inputArray[i+1]
		if maxSolution < product {
			maxSolution = product
		}
	}
	return maxSolution
}
