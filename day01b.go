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

func transformTextNumeralsToInts(line string) string {
	// Replace written numerals with integer versions
	// This is some absolutely hack shit, but it works to solve the oneight and eightwo problem
	// (it would be better to just run once over the string and populate a new thing with the digits, but done is done)
	line = strings.Replace(line, "one", "one1one", -1)
	line = strings.Replace(line, "two", "two2two", -1)
	line = strings.Replace(line, "three", "three3three", -1)
	line = strings.Replace(line, "four", "four4four", -1)
	line = strings.Replace(line, "five", "five5five", -1)
	line = strings.Replace(line, "six", "six6six", -1)
	line = strings.Replace(line, "seven", "seven7seven", -1)
	line = strings.Replace(line, "eight", "eight8eight", -1)
	line = strings.Replace(line, "nine", "nine9nine", -1)
	

	return (line)
}

func getCalibrationPairFromLine(line string) string	{
	// Calibration Pair is first int and last int in a string even if the int is spelled out
	newLine := transformTextNumeralsToInts(line)

	re := regexp.MustCompile(`\d+`)
	vals := re.FindAllString(newLine, -1)

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
	// Digits can be written in text as `one`, `two`, etc
	var inData []byte = getData()
	inDataSplit := strings.Split(string(inData), "\n")

	var runningTotal int64 = 0

	for _, line := range inDataSplit {
			var calibrationVal int64 = getCalibrationValueFromLine(line)
			
			runningTotal += calibrationVal
			fmt.Println(line, " ", calibrationVal, " ", runningTotal)

	}
	fmt.Println("Calibration Value: ", runningTotal)
}