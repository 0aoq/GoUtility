package main

import (
	"fmt"
	"strings"
)

func print(data interface{})                { fmt.Println(data) }            // simple print() function; same as fmt.Println()
func toString(nonstring interface{}) string { return fmt.Sprint(nonstring) } // convert interface{} to string using fmt.Sprint()
func forEach(array []string, callback func(item interface{})) { // same as just a for loop, but simple function
	for _, item := range array {
		callback(item)
	}
}

func main() {
	// test
	str_split := strings.Fields("a b c d")
	forEach(str_split, func(item interface{}) {
		print(toString(item))
	})
}
