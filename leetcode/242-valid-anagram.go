package main

/*
	description: https://leetcode.com/problems/valid-anagram/description/
*/

func isAnagram(s string, t string) bool {

	originMap := make(map[byte]int)
	sByte := []byte(s)

	for _, val := range sByte {
		originMap[val]++
	}

	tByte := []byte(t)
	for _, val := range tByte {
		originMap[val]--
	}

	for _, v := range originMap {
		if v != 0 {
			return false
		}
	}
	return true
}
