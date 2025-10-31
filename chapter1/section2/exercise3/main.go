package main

import "strings"

func main() {

}

func ConcatStrings(times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += "Golang"
	}
	return result
}

func BuildStrings(times int) string {
	var result strings.Builder
	for i := 0; i < times; i++ {
		result.WriteString("Golang")
	}
	return result.String()
}
