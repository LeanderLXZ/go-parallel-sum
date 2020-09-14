/**
* @Author: huadong.hu@outlook.com
* @Date: 8/10/20 10:20
* @Desc:
 */
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Read a list of integers from `fileName`
// and launch `goRoutineNums` go routines to do calculation
// return sum of these Integers
// You Must start exact `goRoutineNums` go routines or you lose points here
func SumRoutine(nums []int, c chan int) {
	// Calculate sum for each go routine
	sum := 0
	for _, v := range nums {
		sum += v
	}
	c <- sum
}

func Sum(goRoutineNum int, fileName string) int {
	//TODO Add your code here
	if goRoutineNum < 1 {
		log.Println("goRoutineNum must >= 1!")
	}

	nums, err := readInts(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Create channel with size of go routine number
	ch := make(chan int, goRoutineNum)

	// Calculate chunk size for each go routine
	var chunkSize int
	if len(nums)%goRoutineNum == 0 {
		chunkSize = len(nums) / goRoutineNum
	} else {
		chunkSize = len(nums)/goRoutineNum + 1
	}

	// Create go routines
	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		go SumRoutine(nums[i:end], ch)
	}

	// Receive value from the channel and calculate the sum of go routines
	sum := 0
	for i := 0; i < goRoutineNum; i++ {
		sum += <-ch
	}
	return sum
}

//Read integers from file
//Do not modify this function
func readInts(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cin := bufio.NewScanner(file)
	cin.Split(bufio.ScanWords)
	var res []int
	for cin.Scan() {
		val, err := strconv.Atoi(cin.Text())
		if err != nil {
			return res, err
		}
		res = append(res, val)
	}
	return res, nil
}
