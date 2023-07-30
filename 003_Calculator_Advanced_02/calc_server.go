package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
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

func (calc *Calculation) printCalculation(conn net.Conn) {
	fmt.Fprintf(conn, "%s = %v\n", calc.cmd, calc.res)
}

func (root *CalcHistory) printFwdHistory(conn net.Conn) {
	currNode := root
	fmt.Fprintln(conn, "---------------------------------")

	for i := 0; ; i++ {
		fmt.Fprintf(conn, "Calc %d: ", i)
		currNode.calc.printCalculation(conn)
		currNode = currNode.nextNode

		if currNode == nil {
			break
		}
	}
	fmt.Fprintln(conn, "---------------------------------")
}

func (lastNode *CalcHistory) printBwdHistory(conn net.Conn) {
	currNode := lastNode
	fmt.Fprintln(conn, "---------------------------------")

	for i := 0; ; i++ {
		fmt.Fprintf(conn, "Calc %d: ", i)
		currNode.calc.printCalculation(conn)
		currNode = currNode.prevNode

		if currNode == nil {
			break
		}
	}
	fmt.Fprintln(conn, "---------------------------------")
}

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var sym string
	var operands []string
	var res float64

	root := &CalcHistory{nil, nil, nil}
	currNode := root

	fmt.Fprintln(conn, "Calculator: An Advanced Take v2.0")
	fmt.Fprintln(conn, "---------------------------------")

	scanner := bufio.NewScanner(conn)

	defer conn.Close()

	for scanner.Scan() {
		cmd := scanner.Text()

		// if cmd == "exit" {
		// 	return
		// }

		operands = strings.Split(cmd, " ")
		fmt.Fprintf(conn, "Your cmd: %v\n", operands)

		if len(operands) == 3 {
			// fmt.Fprintf(conn, "Entering Len 3 Region\n")
			fmt.Fprintf(conn, "")
			if currNode.calc != nil {
				currNode.nextNode = &CalcHistory{}
				currNode.nextNode.prevNode = currNode
				currNode = currNode.nextNode
			}

			n1, err := strconv.ParseFloat(strings.TrimSpace(operands[0]), 64)
			if err != nil {
				fmt.Fprintln(conn, "Parsing: N1 failed")
			}

			sym = strings.TrimSpace(operands[1])

			n2, err := strconv.ParseFloat(strings.TrimSpace(operands[2]), 64)
			if err != nil {
				fmt.Fprintln(conn, "Parsing: N2 failed")
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
					fmt.Fprintf(conn, "Result: Division by zero error\n")
					res = math.Inf(-1)
					break
				}
			}

			// Add this calc to the history
			currNode.calc = &Calculation{strings.TrimSpace(cmd), res}

			fmt.Fprintf(conn, "Result: %v\n", res)
		} else if len(operands) == 2 {
			// fmt.Fprintf(conn, "Entering Len 2 Region\n")
			fmt.Fprintf(conn, "")
			// Process other commands
			if strings.ToLower(strings.TrimSpace(operands[1])) == "history" {
				if strings.ToLower(strings.TrimSpace(operands[0])) == "forward" {
					// fmt.Fprintf(conn, "Entering Forward Region\n")
					fmt.Fprintf(conn, "")
					root.printFwdHistory(conn)
				} else if strings.ToLower(strings.TrimSpace(operands[0])) == "backward" {
					// fmt.Fprintf(conn, "Entering Backward Region\n")
					fmt.Fprintf(conn, "")
					currNode.printBwdHistory(conn)
				}
			}

		} else if len(operands) == 1 {
			// fmt.Fprintf(conn, "Entering Len 1 Region\n")
			fmt.Fprintf(conn, "")
			if strings.ToLower(strings.TrimSpace(operands[0])) == "exit" {
				fmt.Fprintln(conn, "---------------------------------")
				fmt.Fprintln(conn, "Calculator is shutting down...")
				fmt.Fprintln(conn, "Ciao!")
				// os.Exit(0)
				return
			} else {
				fmt.Fprintf(conn, "Error: unknown cmd %s", strings.TrimSpace(operands[0]))
			}
		}

	}
}
