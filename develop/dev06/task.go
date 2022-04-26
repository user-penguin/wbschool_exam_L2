package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"wbschool_exam_L2/develop/dev06/pkg"
)

func main() {
	// задаем флаги
	core := pkg.NewCore()
	flag.IntVar(&core.Fields, "f", 0, "'fields' - выбрать поля (колонки)")
	flag.StringVar(&core.Delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&core.Separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()
	args := flag.Args()
	// байт в условии нет, а значит нужен  f>0, стандартный cut тоже не взлетит
	if core.Fields == 0 {
		log.Fatalln("you must use -f with some value > 0")
	}

	// если файл не добавлен
	if len(args) == 0 {
		terminalReading(core)
	}

	// если есть файл, читаем его
	fileName := args[len(args)-1]
	file, err := ioutil.ReadFile(fileName) // читаем файл
	// если не удалось прочитать, то ожидаем ввод
	if err != nil {
		log.Fatalln(err)
	}

	splitString := strings.Split(string(file), "\n")
	// проходим по всем строкам и для каждой вызываем метод Cut
	for _, str := range splitString {
		if res, ok := Cut(str, core); ok {
			fmt.Println(res)
		}
	}
}

//terminalReading читаем с консоли до закрытия консоли или ctrl-c
func terminalReading(c *pkg.Core) {
	for {
		reader := bufio.NewReader(os.Stdin) //читаем с ввода
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		// "вырезаем" по указанной колонке
		res, _ := Cut(text, c)
		fmt.Println(res)
	}
}

func Cut(str string, c *pkg.Core) (string, bool) {
	// если добавлено -s - строки без delimiter скипаются
	if c.Separated && !strings.Contains(str, c.Delimiter) {
		return "", false
	}
	// пилим строку по Delimiter
	splitStr := strings.Split(str, c.Delimiter)
	if c.Fields <= len(splitStr) {
		return splitStr[c.Fields-1], true
	}
	return "", false
}
