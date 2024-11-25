package main

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
}

func parseBoolExpr(expression string) bool {
	var operands []int32
	var operators []int32
	var result bool
	for _, char := range expression {
		switch char {
		case '&':
			operators = append(operators, char)
		case '|':
			operators = append(operators, char)
		case '!':
			operators = append(operators, char)
		case 't':
			operands = append(operands, char)
		case 'f':
			operands = append(operands, char)
		case '(':
			operands = append(operands, char)
		case ')':
			var value bool
			for i := len(operands) - 1; i >= 0; i-- {
				operand := operands[i]
				if operand == '(' {
					operands = operands[:i]
					break
				}
				if i == len(operands)-1 && operand == 't' {
					value = true
				}
				parse(&value, operand, operators[len(operators)-1])
			}
			if value {
				operands = append(operands, 't')
			} else {
				operands = append(operands, 'f')
			}
			operators = operators[:len(operators)-1]
		}
		var value bool
		for i, operand := range operands {
			if operand == '(' {
				break
			}
			if i == 0 && operand == 't' {
				value = true
			}
			if len(operators) == 0 {
				break
			}
			parse(&value, operand, operators[len(operators)-1])
		}
		result = value
	}
	return result
}

func parse(target *bool, operand, operation int32) {
	var value bool
	if operand == 't' {
		value = true
	}
	switch operation {
	case '|':
		*target = *target || value
	case '&':
		*target = *target && value
	case '!':
		*target = !value
	}
}

func parseBoolExpr2(expression string) bool {
	stack := []rune{}

	for _, c := range expression {
		if c == ')' {
			subExpr := []rune{}
			for stack[len(stack)-1] != '(' {
				subExpr = append(subExpr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // Remove '('

			op := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Remove operator

			if op == '!' {
				if subExpr[0] == 't' {
					stack = append(stack, 'f')
				} else {
					stack = append(stack, 't')
				}
			} else if op == '&' {
				result := 't'
				for _, e := range subExpr {
					if e == 'f' {
						result = 'f'
						break
					}
				}
				stack = append(stack, result)
			} else if op == '|' {
				result := 'f'
				for _, e := range subExpr {
					if e == 't' {
						result = 't'
						break
					}
				}
				stack = append(stack, result)
			}
		} else if c != ',' {
			stack = append(stack, c)
		}
	}

	return stack[0] == 't'
}
