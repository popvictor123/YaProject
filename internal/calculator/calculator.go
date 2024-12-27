package calculator
import (
        "strings"
        "fmt"
        "strconv"
)
func IsNum(s string) bool {
        for _, c := range s {
                if !(48 <= c && c <= 57 || c >= 40 && c <= 43 || c == 45 || c == 47 || c == 46) {
                        return false
                }
        }
        return true
}
func Calc(expression string) (float64, error) {
        var num float64
        var j int
        if !IsNum(expression) {
                return 0, fmt.Errorf("Only numbers and arithmetic operations are allowed")
        }
        i := strings.IndexAny(expression, "+-*/")
        if i < 0 || i == 0 && expression[0] == '-' {
                n, err := strconv.ParseFloat(expression, 64)
                return n, err
        }
        j = strings.Index(expression, "(")
        for j != -1 {
                var cnt, i int = 1, 0
                var val rune
                for i, val = range expression[j + 1:] {
                        if val == '(' {
                                cnt++
                        } else if val == ')' {
                                cnt--
                        }
                        if cnt == 0 {
                                i += j + 1
                                break
                        }
                        if cnt < 0 {
                                return 0, fmt.Errorf("Not right parentheses")
                        }
                }
		if cnt != 0 {
			return 0, fmt.Errorf("Not right parentheses")
		}
		n, err := Calc(expression[j + 1:i])
                if err != nil {
                        return 0, err
                }
                expression = expression[:j] + strconv.FormatFloat(n, 'f', -1, 64) + expression[i + 1:]
                j = strings.Index(expression, "(")

        }

        j = strings.IndexAny(expression, "*/")
        for j != -1 {
                prev := strings.LastIndexAny(expression[:j], "+-*/") + 1
		var next int
		if expression[j + 1] == '-' {
			next = strings.IndexAny(expression[j+2:], "+-/*")
		} else {
			next = strings.IndexAny(expression[j + 1:], "+-/*") 
		}
		if next == -1 {
                        next = len(expression)
                } else if next == 0 {
			return 0, fmt.Errorf("Not enough values")
		}else {
                        next += len(expression[:j + 1])
                }
                var num1, num2 float64
                var err error
                num1, err = strconv.ParseFloat(expression[prev:j], 64)
                if err != nil {
			fmt.Println(expression)
                        return 0, err
                }
                num2, err = strconv.ParseFloat(expression[j + 1:next], 64)
                if err != nil {
			fmt.Println(expression)
                        return 0, err
                }
                if expression[j] == '*' {
                       num = num1 * num2
                } else {
			if num2 == 0 {
				return 0, fmt.Errorf("Division by zero")
			}
                        num = num1 / num2
                }
                n := strconv.FormatFloat(num, 'f', -1, 64)
                expression = expression[:prev] + n + expression[next:]
                j = strings.IndexAny(expression, "*/")
        }
        j = strings.IndexAny(expression[1:], "+-")
        for j != -1 {
                j += 1
                prev := strings.LastIndexAny(expression[:j], "+-*/") + 1
                if expression[0] == '-' {
                        prev = 0
                }
		var next int
		if expression[j + 1] == '-' {
			next = strings.IndexAny(expression[j + 2:], "+-/*") 
		} else {
			next = strings.IndexAny(expression[j + 1:], "+-/*") 
		}
		if next == -1 {
                        next = len(expression)
                } else if next == 0 {
			return 0, fmt.Errorf("Not enough values")
		} else {
                        next += len(expression[:j + 1])
                }
                var num1, num2 float64
                var err error
                num1, err = strconv.ParseFloat(expression[prev:j], 64)
                if err != nil {
                        return 0, err
                }
                num2, err = strconv.ParseFloat(expression[j + 1:next], 64)
                if err != nil {
                        return 0, err
                }
                if expression[j] == '+' {
                       num = num1 + num2
                } else {
                        num = num1 - num2
                }
                n := strconv.FormatFloat(num, 'f', -1, 64)
                expression = expression[:prev] + n + expression[next:]
                j = strings.IndexAny(expression[1:], "+-")
        }
 
        return num, nil
}

