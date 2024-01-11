package main

import "fmt"

type Letters [26]byte

func cntLetters(str string) (letters Letters) {
	for i := range str {
		letters[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Letters][]string)

	for _, v := range strs {
		key := cntLetters(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}