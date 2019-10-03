/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// 主要流程分为：
// 1. 解析命令行 获取文件名文件个数等参数
// 2. 解析文件 对于每个文件进行解析 按行解析 去除非字母 去除过短单词
// 3. 格式化输出
// 具体细节：
// 1. 如何适配unix和windows文件系统
// 2. 如何判断非字母的单词，如何判断单词的长度
// 3. 对于文件EOF的处理
func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	frequencyForWord := map[string]int{}
	for _, filename := range commandLineFiles(os.Args[1:]) {
		fmt.Println(filename)
		updateFrequencies(filename, frequencyForWord)
	}
	reportByWords(frequencyForWord)
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	reportByFrequency(wordsForFrequency)
}

func commandLineFiles(files []string) []string {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			// filepath.Glob(name) 返回一个包含所有文件名符合传入pattern（/usr/*/bin/ed (assuming the Separator is '/')）的切片和错误信息
			fmt.Println(filepath.Glob(name))
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	// 打开文件
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file:", err)
		return
	}
	// 保证文件的关闭
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}

func readAndUpdateFrequencies(reader *bufio.Reader,
	frequencyForWord map[string]int) {
	for {
		// 一行一行的读取文件
		line, err := reader.ReadString('\n')
		// 处理一行中每个单词
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			// 判断只有当长度大于1的才参与统计
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file:", err)
			}
			break
		}
	}
}

// 用来忽略不是单词的字符
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	// strings.FieldsFunc 针对入参s进行切割 返回一个slice，切点在传入函数参数为true的地方
	return strings.FieldsFunc(s, notALetter)
}

func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}

func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency -> Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ","))
	}
}
