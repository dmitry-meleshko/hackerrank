package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY num
 *  2. INTEGER k
 */
func minSum(num []int32, k int32) int32 {
	// Write your code here
	//return minSumDumb(num, k)
	//return minSumInsert(num, k)
	return minSumBacktrack(num, k)
}

// using custom sort function - shouldn't import "sort"?
func insertSort(num []int32, size int32) int64 {
	var key int32
	S := int64(num[0])
	for i := int32(1); i < size; i++ {
		key = num[i]
		S += int64(key)
		j := i - 1
		for j >= 0 && num[j] < key {
			num[j+1] = num[j]
			j--
		}
		num[j+1] = key
	}

	return S
}

// "math" library can't be imported - custom math.Ceil function
func ceil(n float64) int32 {
	var iNum = int32(n)
	if n == float64(iNum) {
		return iNum
	}
	return iNum + 1
}

// minimize backtracking through unsorted part
func minSumBacktrack(num []int32, k int32) int32 {
	// custom sort instead of importing library
	S := insertSort(num, int32(len(num)))

	ixSorted, ixMax := 0, 0 // indicies for Max and Sorted elements
	for n := 0; n < int(k); n++ {
		// divide the max element and perform ceiling
		temp := int32(ceil(float64(num[ixMax]) / 2))
		// reduce running sum by the difference
		S -= int64(num[ixMax] - temp)
		// insert divided number back
		num[ixMax] = temp

		if ixSorted+1 == len(num) {
			//fmt.Println("At the end:", num[:5], "...", num[len(num)-5:])
			S = insertSort(num, int32(len(num)))
			ixMax, ixSorted = 0, 0
			continue
		} else if num[ixSorted+1] > num[ixSorted] {
			//fmt.Println("Move sorted:", ixSorted, ixSorted+1)
			ixSorted++
		}

		// find the largest left-most value
		for i := 0; i <= ixSorted; i++ {
			ixMax = i
			if num[i] >= num[ixSorted] {
				break
			}
		}
	}
	fmt.Println("Sum:", S)
	return int32(S)
}

// Sorts array, afterwards insert element inorder
func minSumInsert(num []int32, k int32) int32 {
	// custom sort instead of importing library
	S := insertSort(num, int32(len(num)))

	// iterate across K divisions - can be longer than array
	n := 0
	for n < int(k) {
		if len(num) > 1 && num[0] < num[1] { // array out of order
			// where does this number fit in array?
			found := false
			temp := make([]int32, len(num))
			for i := 1; i < len(num); i++ {
				if num[0] >= num[i] {
					// prefix - 1st
					copy(temp, num[1:i])
					// insert the element
					copy(temp[i-1:], num[0:1])
					// remainder
					copy(temp[i:], num[i:])
					found = true
					break
				}
			}
			if !found { // all elements are greater
				copy(temp, num[1:])
				copy(temp[len(num)-1:], num[0:1])
			}
			num = temp
		}
		// 1st element should be largest
		temp := int32(ceil(float64(num[0]) / 2))
		// reduce running sum by the difference
		S -= int64(num[0] - temp)
		// insert divided number back
		num[0] = temp
		//fmt.Println(num)
		n++
	}

	// calculate the new sum
	fmt.Println("Sum:", S)
	return int32(S)
}

// Works, but slow -- O(k*len(num))
func minSumDumb(num []int32, k int32) int32 {
	var S, ix int32
	for i := 0; i < int(k); i++ {
		// find maximum
		var max int32
		ix = 0
		for j := 0; j < len(num); j++ {
			if num[j] >= max {
				max = num[j]
				ix = int32(j) // preserve the index
			}
		}
		temp := float64(max) / 2
		// insert divided number back
		num[ix] = int32(ceil(temp))
	}

	// calculate the new sum
	for _, v := range num {
		S += int32(v)
	}

	fmt.Println("Sum:", S)
	return S
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	numCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var num []int32

	for i := 0; i < int(numCount); i++ {
		numItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		numItem := int32(numItemTemp)
		num = append(num, numItem)
	}

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := minSum(num, k)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
