/**
* @Author: huadong.hu@outlook.com
* @Date: 8/17/20 02:09
* @Desc:
 */
package main

import (
	"flag"
	"fmt"
)

func main() {
	sum := 0
	//@TODO read file name and goroutine numbers from command line and output its sum
	//the argument should be `-f` `-g`
	fileName := flag.String("f", "", "file name")
	goRoutines := flag.Int("g", 1, "number of goroutines")
	flag.Parse()
	sum = Sum(*goRoutines, *fileName)

	//DO NOT OUTPUT ANYTHING ABOVE THIS LINE
	//DO NOT MODIFY OUTPUT FORMAT!!
	fmt.Println(sum)
}
