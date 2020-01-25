package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'findSmallestDivisor' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 */
func findSmallestDivisor(s string, t string) int32 {
	// Write your code here

	count := 0 // how many times t is repeated in s
	for {
		// starting index = string length * number of repeats
		start := len(t) * count
		// find first index of the substring
		ix := strings.Index(s[start:], t)
		//fmt.Println(ix)
		if ix == -1 {
			break
		}
		count++
	}

	// no repeats of t or repeats fon't complete the string s
	if count == 0 || strings.Repeat(t, count) != s {
		fmt.Println("Not divisible")
		return -1
	}

	// evaluate upto half of the string
	// grow letter one at a time
	for i := 1; i <= len(t)/2; i++ {
		// 2 letters from 6 letter word are repeated 6/2 times
		if strings.Repeat(t[0:i], len(t)/i) == t {
			fmt.Println("Length:", i)
			return int32(i)
		}
	}
	// only full length word can match t
	fmt.Println("Length:", len(t))

	return int32(len(t))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	result := findSmallestDivisor(s, t)

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
