package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ans_s := make(map[rune]int, 0)
	ans_t := make(map[rune]int, 0)
	for _, val := range s {
		ans_s[val]++
	}
	for _, val := range t {
		ans_t[val]++
	}
	for key, _ := range ans_s {
		if ans_s[key] != ans_t[key] {
			return false
		}
	}
	return true
}
