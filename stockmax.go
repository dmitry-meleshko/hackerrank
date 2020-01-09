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
 * Complete the 'stockmax' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY prices as parameter.
 */

func stockmax(prices []int32) int64 {
	// Write your code here
	if len(prices) == 0 {
		return 0
	}

	var spent, profit, shares int64

	var maxIdx []int // list of indicies for maximum prices
	//fmt.Println("== Prices:", prices)

	var max int32
	ix := 0
	for j := ix; j < len(prices)-1; j++ {
		for i := j; i < len(prices); i++ {
			if prices[i] >= max {
				max = prices[i]
				ix = i // move index forward
			}
		}
		//fmt.Println("ix:", ix)
		maxIdx = append(maxIdx, ix) // save local maximum
		j = ix
		max = 0 // reset
	}

	//fmt.Println("Maximums:", maxIdx)
	ix = 0
	for _, m := range maxIdx {
		lastPrice := prices[ix] // reset price to 1st item in the list
		for i := ix + 1; i <= m; i++ {
			shares++ // buy 1 share per "up" day
			spent += int64(lastPrice)
			//fmt.Println("Shares:", shares)
			//fmt.Println("Spent:", spent)
			lastPrice = prices[i] // today's price
			ix = i                // advance last index
		}
		// sell on the max day
		profit += int64(lastPrice)*shares - spent
		shares = 0
		//fmt.Println("Shares:", shares)
		//fmt.Println("Running Profit:", profit)
		ix++ // reset last index
		spent = 0
	}

	return profit
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		//fmt.Println(n)

		pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var prices []int32

		for i := 0; i < int(n); i++ {
			pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
			checkError(err)
			pricesItem := int32(pricesItemTemp)
			prices = append(prices, pricesItem)
		}

		fmt.Printf("prices: %v: ", len(prices))
		result := stockmax(prices)
		fmt.Println(result)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
