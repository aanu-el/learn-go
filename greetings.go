package main

import "strings"

func sayGreeting(name string) string {
	nameToUpper := strings.ToUpper(name)
	return "Hello, " + nameToUpper + "!"
}