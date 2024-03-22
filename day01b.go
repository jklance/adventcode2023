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
	// TODO: resolve the oneight problem (should return one and eight)
	// TODO: remove oneight from the end of the data file
	replacer := strings.NewReplacer("one", "1one", "two", "2two", "three", "3three", "four", "4four", "five", "5five", "six", "6six", "seven", "7seven", "eight", "8eight", "nine", "9nine")
	newLine := replacer.Replace(line)
	// fmt.Println(newLine)

	return (newLine)
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