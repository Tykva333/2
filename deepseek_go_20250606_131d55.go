package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для упорядочивания чисел по возрастанию
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

// Функция для чтения и обработки файла
func processFile(inputPath, outputPath string) error {
	// Открываем файл для чтения
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var numbers []int64
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Читаем и парсим числа из файла
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		num, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Printf("Пропущено нечисловое значение: %s\n", text)
			continue
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	// Сортируем числа
	sortNumbers(numbers)

	// Создаем файл для записи
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer outFile.Close()

	// Записываем отсортированные числа
	writer := bufio.NewWriter(outFile)
	for _, num := range numbers {
		fmt.Fprintf(writer, "%d ", num)
	}
	writer.Flush()

	fmt.Println("Результат успешно записан в файл")
	fmt.Println("Отсортированные числа:", numbers)
	return nil
}

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	if err := processFile(inputFile, outputFile); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
}