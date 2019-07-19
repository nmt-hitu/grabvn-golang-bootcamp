package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "errors"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan(){
		text := scanner.Text()
		if text != "" {
			fmt.Print(process(text))
			}
		fmt.Print("\n >")
	}

}

func process(text string) string{

	expr := strings.Split(text, " ");

 	a, error := strconv.ParseFloat(expr[0], 10);
	if error != nil {
		return "a has to a number"
	}
	
	b, error := strconv.ParseFloat(expr[2], 10);

	if error !=nil {
		return "b has to a number"
	}	

	op := expr[1];

	var result float64

	switch op {
		case "+":
			result = a + b;
		case "-":
			result = a - b;
		case "*":
			result = a * b;
		case "/":
			if b == 0 {
				return "Invalid value when calculate"
			};
			if b > 0 {
				result = a / b;
			};
		default:
			fmt.Println("Invalid operator.");
	}

	return "Result: " + strconv.FormatFloat(result, 'f', 6, 64);
}