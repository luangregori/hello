package main

import (
	"fmt"
	"regexp"
)

func print(input string, output bool) {
	fmt.Println("Entrada: ", input)
	fmt.Println("Saida: ", output)
	fmt.Println("-----")
}

const pattern_URL string = `^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(((([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)\.)?)?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?)|localhost)(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`
const pattern_HTML string = `<(/|!)?(\w*(\s*)\w*)>`
const pattern_duplicated_string string = `\b(\w+)(\W+\b)+`
const pattern_credit_card string = `^(0[1-9]|1[0-2])(\/|-)([0-9]{4})$`
const pattern_CEP string = `^(\d{0,5}|\d{5}\-\d{0,3})$`

func main() {
	case1 := "http://golang.fastrl.com"
	match1, _ := regexp.MatchString(pattern_URL, case1)
	print(case1, match1)
	case11 := "http::/golang.fastrl.com"
	match11, _ := regexp.MatchString(pattern_URL, case11)
	print(case11, match11)

	re2, _ := regexp.Compile(pattern_HTML)
	case2 := "<html><body></body></html>"
	match2 := re2.MatchString(case2)
	print(case2, match2)
	re22, _ := regexp.Compile(pattern_HTML)
	case22 := "<html"
	match22 := re22.MatchString(case22)
	print(case22, match22)

	re3, _ := regexp.Compile(pattern_duplicated_string)
	case3 := "Lorem ipsum dolor dolor sit amet, consectetur adipiscing elit elit, sed do eiusmod tempor incididunt ut labore labore et dolore magna aliqua."
	match3 := re3.MatchString(case3)
	print(case3, match3)
	re33, _ := regexp.Compile(pattern_duplicated_string)
	case33 := "Lorem"
	match33 := re33.MatchString(case33)
	print(case33, match33)

	re4, _ := regexp.Compile(pattern_credit_card)
	case4 := "02/2413"
	match4 := re4.MatchString(case4)
	print(case4, match4)
	re44, _ := regexp.Compile(pattern_credit_card)
	case44 := "22/2413"
	match44 := re44.MatchString(case44)
	print(case44, match44)

	re5, _ := regexp.Compile(pattern_CEP)
	case5 := "99950-000"
	match5 := re5.MatchString(case5)
	print(case5, match5)
	re55, _ := regexp.Compile(pattern_CEP)
	case55 := "99950000"
	match55 := re55.MatchString(case55)
	print(case55, match55)
}
