package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getData() []byte {
	inData, err := os.ReadFile("day01a-data.txt")
	if(err != nil) {
		panic(err)
	}
	return (inData)
}

/*func getFirstIntFromString(line string) {
	if(!(len(line) > 0)) {
		panic("String too short")
	}
	firstInt := line[0][0:1]
	return (firstInt)
}*/

func main() {
	// Combine the first and last digit from each line to form a value, get running total of all
	var inData []byte = getData()

	inDataSplit := strings.Split(string(inData), "\n")
	var runningTotal int64 = 0

	for _, line := range inDataSplit {
			re := regexp.MustCompile(`\d+`)
			vals := re.FindAllString(line, -1)

			var first string = vals[0][0:1]
			var last string = vals[len(vals) - 1]
			last = last[len(last)-1:]

			var lineVal string = first + last
		
			calibrationVal, err := strconv.ParseInt(lineVal, 0, 64)
			if(err != nil) {
				panic(err)
			}
			runningTotal += calibrationVal
			fmt.Println(line, " ", lineVal, " ", calibrationVal, " ", runningTotal)

	}
	fmt.Println(runningTotal)
}