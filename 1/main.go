package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	hashmap := make(map[int][]int)
	file, err := os.Open("./note.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	line_count := 1
	for scanner.Scan() {
		line := scanner.Text()
		// println(line)

		re := regexp.MustCompile("[0-9]+")
		numberStr := re.FindAllString(line, -1)

		for _, v := range numberStr {
			new, _ := strconv.Atoi(v)
			hashmap[line_count] = append(hashmap[line_count], new)
		}
		line_count++
	}

	sum := 0
	for _, val := range hashmap {

		if len(val) == 1 {
			for _, v := range val {
				n_str := strconv.Itoa(v)

				if len(n_str) > 1 {
					first := string(n_str[0])
					last := string(n_str[len(n_str)-1])
					comb := first + last
					comb_int, _ := strconv.Atoi(comb)
					sum += comb_int
				} else {
					first := n_str + n_str
					first_int, _ := strconv.Atoi(first)
					sum += first_int
				}
			}
		} else {
			str := ""
			for _, v := range val {
				str += strconv.Itoa(v)
			}
			first := string(str[0])
			last := string(str[len(str)-1])
			comb := first + last
			comb_int, _ := strconv.Atoi(comb)
			sum += comb_int
		}
	}
	println(sum)
}
