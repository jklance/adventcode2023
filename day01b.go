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

func getFirstIntFromString(line string) string {
	if(!(len(line) > 0)) {
		panic("String too short")
	}
	firstInt := line[0:1]
	return (firstInt)
}

func getLastIntFromString(line string) string {
	if(!(len(line) > 0)) {
		panic("String too short")
	}
	lastInt := line[len(line) - 1:]
	return (lastInt)
}

func getCalibrationPairFromLine(line string) string{
	// Calibration Pair is first int and last int in a string
	re := regexp.MustCompile(`\d+`)
	vals := re.FindAllString(line, -1)

	first := getFirstIntFromString(vals[0])
	last := getLastIntFromString(vals[len(vals) - 1])
	calibrationPair := first + last

	return (calibrationPair)
}

func getCalibrationValueFromLine(line string) int64 {
	// Calibration Value is the first and last integer in a line put together to form a 2-digit number
	calibrationPair := getCalibrationPairFromLine(line)
					
	calibrationVal, err := strconv.ParseInt(calibrationPair, 0, 64)
	if(err != nil) {
		panic(err)
	}

	return (calibrationVal)
}


func main() {
	// Combine the first and last digit from each line to form a value, get running total of all
	var inData []byte = getData()
	inDataSplit := strings.Split(string(inData), "\n")

	var runningTotal int64 = 0

	for _, line := range inDataSplit {
			var calibrationVal int64 = getCalibrationValueFromLine(line)
			
			runningTotal += calibrationVal
			// fmt.Println(line, " ", calibrationVal, " ", runningTotal)

	}
	fmt.Println("Calibration Value: ", runningTotal)
}