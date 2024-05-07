package array_and_hashing

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


/*
optimize
1, use char slice which its length is 26, because only 26 lower letter in the question.
2, use rune to traverse the string if this question require utf8mb4 character.
*/