package main

func solution(inputString string) bool {
	for i := 0; i < len(inputString); i++ {
		if inputString[i] != inputString[len(inputString)-i-1] {
			return false
		}
	}
	return true
}
