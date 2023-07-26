package main

import (
	"bufio"
	"fmt"
	"math"
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
	calc     *Calculation
	nextNode *CalcHistory
	prevNode *CalcHistory
}

func (calc *Calculation) printCalculation() {
	fmt.Printf("%s = %v\n", calc.cmd, calc.res)
}

func (root *CalcHistory) printFwdHistory() {
	currNode := root
	fmt.Println("---------------------------------")

	for i := 0; ; i++ {
		fmt.Printf("Calc %d: ", i)
		currNode.calc.printCalculation()
		currNode = currNode.nextNode

		if currNode == nil {
			break
		}
	}
	fmt.Println("---------------------------------")
}

func (lastNode *CalcHistory) printBwdHistory() {
	currNode := lastNode
	fmt.Println("---------------------------------")

	for i := 0; ; i++ {
		fmt.Printf("Calc %d: ", i)
		currNode.calc.printCalculation()
		currNode = currNode.prevNode

		if currNode == nil {
			break
		}
	}
	fmt.Println("---------------------------------")
}

func main() {
	var sym string
	var operands []string
	var res float64

	root := &CalcHistory{nil, nil, nil}
	currNode := root

	fmt.Println("Calculator: An Advanced Take v1.0")
	fmt.Println("---------------------------------")

	for {
		fmt.Printf("Input cmd: ")
		stdIn := bufio.NewReader(os.Stdin)
		cmd, err := stdIn.ReadString('\n')
		if err != nil {
			fmt.Println("Parsing: input cmd failed")
		}

		operands = strings.Split(cmd, " ")

		if len(operands) == 3 {

			if currNode.calc != nil {
				currNode.nextNode = &CalcHistory{}
				currNode.nextNode.prevNode = currNode
				currNode = currNode.nextNode
			}

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
					res = math.Inf(-1)
					break
				}
			}

			// Add this calc to the history
			currNode.calc = &Calculation{strings.TrimSpace(cmd), res}

			fmt.Printf("Result: %v\n", res)
		} else if len(operands) == 2 {
			// Process other commands
			if strings.ToLower(strings.TrimSpace(operands[1])) == "history" {
				if strings.ToLower(strings.TrimSpace(operands[0])) == "forward" {
					root.printFwdHistory()
				} else if strings.ToLower(strings.TrimSpace(operands[0])) == "backward" {
					currNode.printBwdHistory()
				}
			}

		} else if len(operands) == 1 {
			if strings.ToLower(strings.TrimSpace(operands[0])) == "exit" {
				fmt.Println("---------------------------------")
				fmt.Println("Calculator is shutting down...")
				fmt.Println("Ciao!")
				os.Exit(0)
			} else {
				fmt.Printf("Error: unknown cmd %s", strings.TrimSpace(operands[0]))
			}
		}
	}
}
