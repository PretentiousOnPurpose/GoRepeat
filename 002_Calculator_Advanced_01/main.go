package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Capture calculation history via LinkedLists

type Calculation struct {
	cmd string
	res float64
}

type CalcHistory struct {
	calc     Calculation
	nextNode *CalcHistory
	prevNode *CalcHistory
}

func (calc *Calculation) printCalculation() {
	fmt.Printf("%s = %f\n", calc.cmd, calc.res)
}

func (root *CalcHistory) printFwdHistory() {
	currNode := root

	for i := 0; ; i++ {
		fmt.Printf("Calc %d: ", i)
		currNode.calc.printCalculation()
		currNode = currNode.nextNode

		if currNode == nil {
			break
		}
	}
}

func (lastNode *CalcHistory) printBwdHistory() {
	currNode := lastNode

	for i := 0; ; i++ {
		fmt.Printf("Calc %d: ", i)
		currNode.calc.printCalculation()
		currNode = currNode.prevNode

		if currNode == nil {
			break
		}
	}
}

func main() {
	var sym string
	var operands []string
	var res float64

	fmt.Println("Calculator: An Advanced Take v1.0")
	fmt.Println("------------------------")

	for {
		fmt.Printf("Input cmd: ")
		stdIn := bufio.NewReader(os.Stdin)
		cmd, err := stdIn.ReadString('\n')
		if err != nil {
			fmt.Println("Parsing: input cmd failed")
		}

		operands = strings.Split(cmd, " ")

		if len(operands) == 3 {
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

			// Add this calc to the history

			fmt.Printf("Result: %v\n", res)
		} else {
			// Process other commands
		}
	}
}
