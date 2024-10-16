package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func plusMinus(arr []int32) {
	// Write your code here
	denom := len(arr)
	pos, neg, zero := 0, 0, 0
	for i := range arr {
		if arr[i] == 0 {
			zero++
		} else if arr[i] > 0 {
			pos++
		} else {
			neg++
		}
	}

	slices.Sort(arr)

	fmt.Println(len(arr) / 2)
	fmt.Printf("%.6f", float32(pos)/float32(denom))
	fmt.Printf("%.6f", float32(neg)/float32(denom))
	fmt.Printf("%.6f", float32(zero)/float32(denom))
}

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func miniMaxSum(arr []int32) {
	// Write your code here
	small := slices.Min(arr)
	big := slices.Max(arr)

	var maxSum int64 = 0
	var minSum int64 = 0

	for i := range arr {
		maxSum += int64(arr[i])
		minSum += int64(arr[i])
	}

	maxSum = maxSum - int64(small)
	minSum = minSum - int64(big)

	fmt.Printf("%d %d", minSum, maxSum)
}

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
// AM													PM
// 12 11 10 9 8 7 6 5 4 3 2 1 2 3 4 5 6 7 8 9 10 11 12
// 00 11 10 9 8 7 6 5 4 3 2 13 14 15 16 17 18 19 20 21 22 23 24

func timeConversion(s string) string {
	// Write your code here
	timed := strings.Split(s, ":")
	hour, _ := strconv.Atoi(timed[0])
	min := timed[1]
	seconds := timed[2][:2]
	meri := timed[2][2:]

	if hour == 12 {
		if meri == "AM" {
			hour = 0
		}
	} else if meri == "PM" {
		hour += 12
	}

	return fmt.Sprintf("%02d:%s:%s", hour, min, seconds)
}

func lonelyinteger(a []int32) int32 {
	// Write your code here
	var result int32

	for _, num := range a {
		result ^= num
	}

	return result
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var left int32 = 0
	var right int32 = 0

	for i := 0; i < len(arr); i++ {
		left += arr[i][i]
		right += arr[i][len(arr)-1-i]
	}
	return int32(math.Abs(float64(left - right)))
}

func main() {
	draws, err := GetNumDraws(2011)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of draws in 2011: %d\n", draws)

	goals, err := GetWinnerTotalGoals("UEFA Champions League", 2011)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total goals for the winner of UEFA Champions League 2011: %d\n", goals)
	// matrix := [][]int32{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{9, 8, 9},
	// }
	// out := diagonalDifference(matrix)
	// fmt.Println(out)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
