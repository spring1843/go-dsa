package recursion

func ExpressionOperators(list []int, target int) string {
	if len(list) <= 1 {
		return ""
	}
	return expressionOperatorsHelper(list[1:], list[0], "", target)
}

func expressionOperatorsHelper(list []int, current int, operators string, target int) string {
	if len(list) == 0 {
		if current == target {
			return operators
		}
		return ""
	}

	// Try with the + operator
	result := expressionOperatorsHelper(list[1:], current+list[0], operators+"+", target)
	if result != "" {
		return result
	}

	// Try with the - operator
	result = expressionOperatorsHelper(list[1:], current-list[0], operators+"-", target)
	if result != "" {
		return result
	}

	return ""
}
