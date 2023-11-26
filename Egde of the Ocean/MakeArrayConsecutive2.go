package main

func solution(statues []int) int {
	n := len(statues)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if statues[j] < statues[j+1] {
				swap := statues[j]
				statues[j] = statues[j+1]
				statues[j+1] = swap
			}
		}
	}
	countNeed := 0
	for i := 0; i < len(statues)-1; i++ {
		countNeed += statues[i] - statues[i+1] - 1
	}
	return countNeed
}
