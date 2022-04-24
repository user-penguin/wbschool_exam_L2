/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type IndexValue struct {
	index int
	value string
}

func sortIndexValueArray(src []IndexValue, asc bool) []int {
	index := make(map[string]int)
	var sorting []string
	for _, iv := range src {
		index[iv.value] = iv.index
		sorting = append(sorting, iv.value)
	}
	sort.Strings(sorting)
	var newIndexes []int
	switch asc {
	case true:
		for i := 0; i < len(sorting); i++ {
			newIndexes = append(newIndexes, index[sorting[i]])
		}
	case false:
		for i := len(sorting) - 1; i >= 0; i-- {
			newIndexes = append(newIndexes, index[sorting[i]])
		}
	}

	return newIndexes
}

func main() {
	var column int
	var n, r, u bool
	flag.IntVar(&column, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename == "" {
		fmt.Printf("Не указано имя файла.\n")
		os.Exit(1)
	} else {
		f, err := os.Open(filename)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Printf("Error closing fiel: %s", err)
			}
		}(f)
		if err != nil {
			fmt.Printf("Error opening file: %s", err)
			os.Exit(1)
		}
		in = f
	}

	var lines [][]string
	buf := bufio.NewScanner(in) //таргетим сканнер на файл, который успешно открылся
	for buf.Scan() {
		line := buf.Text()
		lines = append(lines, strings.Split(line, " "))
	}
	var ivs []IndexValue
	for i, str := range lines {
		ivs = append(ivs, IndexValue{
			index: i,
			value: str[column],
		})
	}
	newOrder := sortIndexValueArray(ivs, !r)
	for _, i := range newOrder {
		fmt.Println(lines[i])
	}
}
