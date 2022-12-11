/*
QUESTION:
Write a program which prompts the user to enter a string. The program searches through
the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!”
if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains
the character ‘a’. The program should print “Not Found!” otherwise. The program should not be
case-sensitive, so it does not matter if the characters are upper-case or lower-case.

EXAMPLE: The program should print “Found!” for the following example entered strings,
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for
the following strings, “ihhhhhn”, “ina”, “xian”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter a string=")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	fmt.Printf("\tYour string: \t\"%v\" \n", inputString)

	lowerString := strings.ToLower(inputString)
	fmt.Printf("\t[Step-I] Change to lower case= \t%v \n", lowerString)

	formattedString := strings.ReplaceAll(lowerString, " ", "")
	fmt.Printf("\t[Step-II] Ignore all spaces= \t%v \n", formattedString)
	fmt.Printf("\t[Step-III] Find 'i a n'. . .\t\n")

	fmt.Println("===============================================================")
	if strings.Contains(formattedString, "a") {
		if strings.HasPrefix(formattedString, "i") {
			if strings.HasSuffix(formattedString, "n") {
				fmt.Println("Found!")
				os.Exit(0)
			}
		}
	}
	fmt.Println("Not Found!")
}
