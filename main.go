package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
}

func findCalibrationString() int {
	filePath := "./input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	sum := 0

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		sum = sum + getCalibrationString(line)
		if err != nil {
			break
		}
	}

	return sum
}

func getCalibrationString(line string) int {
	calib := ""
	for _, r := range line {
		if unicode.IsDigit(r) {
			calib = calib + string(r)
		}
	}

	if len(calib) == 1 {
		calib = calib + calib
		calibInt, err := strconv.Atoi(calib)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return 0
		}
		return calibInt
	} else if len(calib) <= 1 {
		return 0
	}

	firstRune := calib[0]
	lastRune := calib[len(calib)-1]
	calib = string(firstRune) + string(lastRune)
	calibInt, err := strconv.Atoi(calib)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return calibInt
}
