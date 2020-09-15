/**
* @Author: huadong.hu@outlook.com
* @Date: 8/10/20 10:20
* @Desc:
 */
package main

import (
	"bufio"

	// "fmt"
	"log"
	"os"
	"strconv"
)

// Read a list of integers from `fileName`
// and launch `goRoutineNums` go routines to do calculation
// return sum of these Integers
// You Must start exact `goRoutineNums` go routines or you lose points here
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
	ch := make(chan int)

	// Create go routines
	for i := 0; i < goRoutineNum; i++ {
		chunk := nums[i*len(nums)/goRoutineNum : (i+1)*len(nums)/goRoutineNum]
		go func() {
			sumRoutine := 0
			for _, v := range chunk {
				sumRoutine += v
			}
			ch <- sumRoutine
		}()
	}

	// Receive value from the channel and calculate the sum of go routines
	sum := 0
	for i := 0; i < goRoutineNum; i++ {
		sum += <-ch
		// Check if the threads are parallel
		// fmt.Println(sum)
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
