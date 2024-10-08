package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//mergeAlternately("abc", "pqr")
	//gcdTryAgain("ABAB", "AB")
	// arr := []int{2, 3, 5, 1, 3}
	// kidsWithCandies(arr, 3)
	// out := canPlaceFlowersAgain([]int{1, 0, 0, 0, 1}, 1)
	// out := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	// out := reverseVowels("IceCreAm")
	// fmt.Println(out)
	out := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
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

func removeDuplicates(nums []int) int {
	// https://gist.github.com/johnwesonga/6301924?permalink_comment_id=4722148#gistcomment-4722148
	// By using a map I can check if I have seen the number before
	// Time: 0(N)?
	// Space: 0(N)

	// seen := make(map[int]bool, len(nums))
	// k := 0
	// for i := 0; i < len(nums); i++ {
	j := 1 //using the j solution: Time 0(1)
	for i := 1; i < len(nums); i++ {
		// if !seen[nums[i]] {
		if nums[j-1] != nums[i] {
			nums[j] = nums[i]
			// seen[nums[i]] = true
			j++
		}
	}
	fmt.Println(nums)
	return j
}

func reverseVowels(s string) string {
	// IceCreAm
	// https://leetcode.com/problems/reverse-vowels-of-a-string/solutions/3934949/detailed-solution-in-go-o-n-time-space-complexity
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for i, v := range s {
		fmt.Printf("%d | Here we have letter: %v\n", i, string(v))
		for _, j := range vowels {
			if strings.ToLower(string(v)) == string(j) {
				fmt.Println(string(j))
			}
		}
	}
	return ""
}

func removeDuplicatesII(nums []int) int {
	myMap := make(map[int]int)
	count := 0

	for i := 0; i < len(nums); i++ {
		if myMap[nums[i]] < 2 {
			myMap[nums[i]] = myMap[nums[i]] + 1
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

func productExceptSelf(nums []int) []int {
	// https://www.youtube.com/watch?v=bNvIQI2wAjk
	ans := make([]int, len(nums))

	for i := range ans {
		ans[i] = 1
	}

	leftOutput := 1
	for i := 0; i < len(nums); i++ {
		ans[i] = leftOutput
		leftOutput *= nums[i]
	}
	rightOutput := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= rightOutput
		rightOutput *= nums[i]
	}

	return ans
}

func majorityElement(nums []int) int {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = myMap[nums[i]] + 1
	}
	fmt.Println(myMap)
	for i := range myMap {
		if myMap[i] > len(nums)/2 {
			return i
		}
	}
	return 1
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // Handle cases where k is greater than n
	if k == 0 {
		return // No need to rotate
	}

	// Reverse helper function
	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	// Step 1: Reverse the entire array
	reverse(0, n-1)
	// Step 2: Reverse the first k elements
	reverse(0, k-1)
	// Step 3: Reverse the remaining elements
	reverse(k, n-1)
}

func increasingTriplet(nums []int) bool {
	inf := math.Pow(10, 20)
	s1 := inf
	s2 := inf
	for i := range nums {
		if float64(nums[i]) > s2 {
			return true
		}
		if float64(nums[i]) > s1 {
			s2 = min(float64(nums[i]), s2)
		}
		s1 = min(float64(nums[i]), s1)
	}
	return false
	// for i := 0; i < len(nums)-2; i++ {
	// 	if nums[i] < nums[i+1] && nums[i+1] < nums[i+2] {
	// 		return true
	// 	}
	// }
	// return false
}

func compress(chars []byte) int {
	writePos := 0
	readPos := 0

	for readPos < len(chars) {
		currentChar := chars[readPos]

		count := 0
		for readPos < len(chars) && chars[readPos] == currentChar {
			readPos++
			count++
		}

		chars[writePos] = currentChar
		writePos++

		if count > 1 {
			countStr := strconv.Itoa(count)

			for i := 0; i < len(countStr); i++ {
				chars[writePos] = countStr[i]
				writePos++
			}
		}
	}

	return writePos
}

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = max(maxProfit, profit)
		} else {
			l = r
		}
		r += 1
	}
	return maxProfit
}
