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
func SumRoutine(nums []int, c chan int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func Sum(goRoutineNum int, fileName string) int {
	//TODO Add your code here
	nums, err := readInts(fileName)
	if err != nil {
		log.Fatal(err)
	}
	chunkSize := len(nums) / goRoutineNum
	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		chunk := nums[i:end]
		go SumRoutine(chunk)
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
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
