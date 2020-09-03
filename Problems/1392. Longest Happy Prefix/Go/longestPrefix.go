func longestPrefix(s string) string {
	arr := make([]int, len(s))
	arr[0] = 0
	for i, j := 1, 0; i < len(s); {
		if s[i] == s[j] {
			arr[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = arr[j - 1]
			}
		}
	}
	return s[0:arr[len(s) - 1]]
}
