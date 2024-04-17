package main

import "fmt"

func slidingWindow(s string) int {
	left, right := 0, 0
	cSet := make(map[byte]int)
	m := 0
	for right < len(s) {
		cSet[s[right]]++
		for cSet[s[right]] > 1 {
			cSet[s[left]]--
			if cSet[s[left]] == 0 {
				delete(cSet, s[left])
			}
			left++
		}
		if len(cSet) > m {
			m = len(cSet)
		}
		right++
	}
	return m
}

func main() {
	fmt.Println(slidingWindow("ttkklmalmm"))
	fmt.Println(slidingWindow("kkllmm"))
}
