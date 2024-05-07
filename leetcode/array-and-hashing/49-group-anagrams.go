package main

/*
description: https://leetcode.com/problems/group-anagrams/

解题思路
1，将字符串分解成字符之后再重新按字母顺序组合成新的字符串
2，用一个 map 来保存拥有相同的字母顺序的字符串
*/

func groupAnagrams(strs []string) [][]string {

	mp := make(map[string][]string)

	for _, strVal := range strs {
		var charArr [26]int32
		for _, val := range strVal {
			charArr[val-'a']++
		}

		str := buildString(charArr)
		_, exist := mp[str]
		if !exist {
			mp[str] = make([]string, 0)
		}
		mp[str] = append(mp[str], strVal)
		//怎么归集？？？
	}

	result := make([][]string, 0)

	for _, v := range mp {
		result = append(result, v)
	}
	return result
}

func buildString(arr [26]int32) string {
	result := make([]rune, 0)
	for k, cnt := range arr {
		for i := 0; i < int(cnt); i++ {
			result = append(result, rune(k)+'a')
		}
	}
	return string(result)
}

/*
//直接定义一个新的类型，一个长度为 26的数组，再做对应的 Map
//time complexity O ( n * L ) n is the number of string in list, L is the length of string
//space complexity O ( n * L )

type Key [26]byte

func groupAnagrams(strs []string) [][]string {
    groups := make(map[Key][]string)
    for _, v := range strs {
        key := strkey(v)
        groups[key] = append(groups[key],v)
    }
    result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

func strkey(str string) (key Key) {
    for i := range str {
		key[str[i]-'a']++
	}
	return key

}
*/
