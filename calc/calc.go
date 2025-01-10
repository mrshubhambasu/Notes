package calc

import "strings"

func Sum(a, b int) (int, []int) {
	s := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		s = append(s, a+b)
	}

	return a + b, s
}

func CountVowels(s string) int {
	volwels := "aeiouAEIOU"
	count := 0

	for _, v := range s {

		if strings.Contains(volwels, string(v)) {
			count++
		}
	}

	return count
}

func CountVowelsOPT(s string) int {
	volwels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	count := 0

	for _, v := range s {
		if volwels[v] {
			count++
		}
	}

	return count
}
func GenerateStr(size int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, size)

	for i := range result {
		result[i] = letters[i%len(letters)]

	}

	return string(result)
}
