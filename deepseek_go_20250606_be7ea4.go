package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sortNumbers(numbers []int64) {
	length := len(numbers)
	for i := 0; i < length-1; i++ {
		swapped := false
		for j := 0; j < length-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func processFile(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var numbers []int64
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
	