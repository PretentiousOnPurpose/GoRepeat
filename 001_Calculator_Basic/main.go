package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sym string
	var operands []string
	var res float64

	fmt.Println("Calculator: A Simple App")
	fmt.Println("------------------------")

	for {
		fmt.Printf("Input cmd: ")
		stdIn := bufio.NewReader(os.Stdin)
		cmd, err := stdIn.ReadString('\n')
		if err != nil {
			fmt.Println("Parsing: input cmd failed")
		}

		operands = strings.Split(cmd, " ")

		n1, err := strconv.ParseFloat(strings.TrimSpace(operands[0]), 64)
		if err != nil {
			fmt.Println("Parsing: N1 failed")
		}

		sym = strings.TrimSpace(operands[1])

		n2, err := strconv.ParseFloat(strings.TrimSpace(operands[2]), 64)
		if err != nil {
			fmt.Println("Parsing: N2 failed")
		}

		switch sym {
		case "+":
			res = n1 + n2
		case "-":
			res = n1 - n2
		case "*":
			res = n1 * n2
		case "/":
			if n2 != 0 {
				res = n1 / n2
			} else {
				fmt.Printf("Result: Division by zero error\n")
				break
			}
		}

		fmt.Printf("Result: %v\n", res)
	}
}
