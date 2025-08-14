package main

import (
	"fmt"
	"strings"
);

func getInitials(n string) (string, string) {
	fullName := strings.ToUpper(n)
	namesArr := strings.Split(fullName, " ")

	var initials []string

	for _, names := range namesArr {
		initials = append(initials, names[:1]) 
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} 

	return initials[0], "_"
}

func main() {
	firstName, lastName := getInitials("Peace")
	fmt.Printf("Your initials are: %s %s \n", firstName, lastName)
}