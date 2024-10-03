package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	//mergeAlternately("abc", "pqr")
	//gcdTryAgain("ABAB", "AB")
	// arr := []int{2, 3, 5, 1, 3}
	// kidsWithCandies(arr, 3)
	// out := canPlaceFlowersAgain([]int{1, 0, 0, 0, 1}, 1)
	out := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(out)
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
	return greatest
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	// By not padding this it was not possible to check the previous number

	for i, current := range flowerbed[:len(flowerbed)-1] {
		if current == 1 && flowerbed[i+1] == 0 {
			fmt.Println(i, flowerbed[i])
			i++
			flowerbed[i+1] = 1
			current = flowerbed[i+1]
			n--
		}
	}
	fmt.Println(flowerbed)
	return false
}

func canPlaceFlowersAgain(flowerbed []int, n int) bool {
	// Space COmplexity: 0{1}
	// TIME Complexity: 0{n}
	// f := slices.Concat([]int{0}, flowerbed, []int{0})
	f := append([]int{0}, flowerbed...)
	f = append(f, 0)
	for i := 1; i < len(f)-1; i++ {
		if f[i-1] == 0 && f[i] == 0 && f[i+1] == 0 {
			f[i] = 1
			n--
		}
	}
	return n <= 0
}

func removeElement(nums []int, val int) int {
	// q
	k := 0
	for i := 0; i < len(nums); i++ {
		// 0, 1, 2, 2, 3, 0, 4, 2 | 2
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(nums)
	return k
}
