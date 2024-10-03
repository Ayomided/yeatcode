package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// out := mergeAlternately("abc", "pqr")
	// out := gcdTryAgain("ABAB", "AB")
	arr := []int{2, 3, 5, 1, 3}
	kidsWithCandies(arr, 3)
	// fmt.Println(out)
}

func mergeAlternately(word1 string, word2 string) string {
	var result strings.Builder
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		result.WriteByte(word1[i])
		result.WriteByte(word2[j])
		i++
		j++
	}

	result.WriteString(word1[i:])
	result.WriteString(word2[j:])

	return result.String()
}

func gcdOfStrings(str1 string, str2 string) string {
	i, j := 0, 0
	x := []string{}

	strarr1 := strings.Split(str1, "")
	strarr2 := strings.Split(str2, "")

	for i < len(strarr1) && j < len(strarr2) {
		if strarr1[i] == strarr2[j] {
			x = append(x, strarr1[i])
		}
		i++
		j++

	}
	return strings.Join(x, "")
}

func gcdTryAgain(str1, str2 string) string {
	// Space Complexity: 0(1)
	// Time Complexity: 0(N)
	len1 := len(str1)
	len2 := len(str2)

	for i := min(len1, len2); i > 0; i-- {
		isDivisor := func(l int) bool {
			if len1%l != 0 || len2%l != 0 {
				return false
			}
			f1, f2 := len1/l, len2/l
			return strings.Repeat(str1[:l], f1) == str1 && strings.Repeat(str1[:l], f2) == str2
		}
		if isDivisor(i) {
			return str1[:i]
		}
	}
	return ""
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var greatest []bool

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= slices.Max(candies) {
			greatest = append(greatest, true)
		} else {
			greatest = append(greatest, false)
		}
	}
	fmt.Println(greatest)
	// return greatest
}
