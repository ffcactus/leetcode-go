package basic_calculator

import (
	"fmt"
	"strconv"
)

const (
	operatorType = 1
	valueType    = 2

	operationAdd        = 1
	operationSub        = 2
	operationLeftBrace  = 3
	operationRightBrace = 4
)

type token struct {
	t         int
	value     int
	operation int
}

func (t token) String() string {
	if t.t == operatorType {
		switch t.operation {
		case operationAdd:
			return "+"
		case operationSub:
			return "-"
		case operationLeftBrace:
			return "("
		case operationRightBrace:
			return ")"
		default:
			return "?"
		}
	}
	return fmt.Sprintf("%d", t.value)
}

func calculate(s string) int {
	tokens, err := parsing(s)
	if err != nil {
		return 0
	}
	return valueOf(tokens)
}

func valueOf(tokens []token) int {
	l := len(tokens)

	if l == 0 {
		return 0
	}
	if l == 1 {
		if tokens[0].t == valueType {
			return tokens[0].value
		} else {
			return 0
		}
	}
	if l == 2 {
		fmt.Println("valueOf 2 tokens")
		return 0
	}
	// 3个或3个以上tokens.

	// 3个tokens是，数字、操作符、数字
	if tokens[0].t == valueType && tokens[1].t == operatorType && tokens[2].t == valueType {
		if l == 3 {
			return cal(tokens[0:1], tokens[1].operation, tokens[2:3])
		}
		if l == 4 {
			fmt.Println("valueOf 4 tokens.")
			return 0
		}
		if tokens[3].t != operatorType {
			fmt.Println("valueOf >4 tokens, token 4 invalid")
			return 0
		}
		tmpToken := token{
			t:     valueType,
			value: cal(tokens[0:1], tokens[1].operation, tokens[2:3]),
		}
		return valueOf(append([]token{tmpToken}, tokens[3:]...))
	}

	// 第一个是（，必须有4个以上的tokens.
	if tokens[0].t == operatorType && tokens[0].operation == operationLeftBrace {
		// 找到与之匹配的‘）’
		i := 1
		moreLeftBrace := 0
		for ; i < l; i++ {
			if tokens[i].t == operatorType && tokens[i].operation == operationRightBrace {
				if moreLeftBrace == 0 {
					break
				} else {
					moreLeftBrace--
				}
			}
			if tokens[i].t == operatorType && tokens[i].operation == operationLeftBrace {
				moreLeftBrace++
			}
		}
		if moreLeftBrace != 0 {
			fmt.Println("valueOf 3 tokens, brace not match.")
			return 0
		}
		if i == l {
			fmt.Println("valueOf 3 tokens, ) not found.")
			return 0
		}
		tempToken := token{
			t:     valueType,
			value: valueOf(tokens[1:i]),
		}
		return valueOf(append([]token{tempToken}, tokens[i+1:]...))
	}

	// num op ( 的形式。
	if tokens[0].t == valueType && tokens[1].t == operatorType && tokens[2].t == operatorType && tokens[2].operation == operationLeftBrace {
		// 找到与之匹配的‘）’
		i := 3
		moreLeftBrace := 0
		for ; i < l; i++ {
			if tokens[i].t == operatorType && tokens[i].operation == operationRightBrace {
				if moreLeftBrace == 0 {
					break
				} else {
					moreLeftBrace--
				}
			}
			if tokens[i].t == operatorType && tokens[i].operation == operationLeftBrace {
				moreLeftBrace++
			}
		}
		if moreLeftBrace != 0 {
			fmt.Println("valueOf 3 tokens, brace not match.")
			return 0
		}
		if i == l {
			fmt.Println("valueOf 3 tokens, ) not found.")
			return 0
		}
		tempToken := token{
			t:     valueType,
			value: valueOf(tokens[3:i]),
		}
		newTokens := tokens[0:2]
		newTokens = append(newTokens, tempToken)
		newTokens = append(newTokens, tokens[i+1:]...)
		return valueOf(newTokens)
	}
	fmt.Println("valueOf >3 tokens, bad format.")
	return 0
}

func cal(left []token, op int, right []token) int {
	l := valueOf(left)
	r := valueOf(right)
	// fmt.Println("cal", left, l, right, r)
	switch op {
	case operationAdd:
		return l + r
	case operationSub:
		return l - r
	}
	fmt.Println("cal unknown op", op)
	return 0
}

func parsing(s string) ([]token, error) {
	ret := make([]token, 0)
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			ret = append(ret, token{
				t:         operatorType,
				operation: operationLeftBrace,
			})
			i++
			continue
		}
		if s[i] == ')' {
			ret = append(ret, token{
				t:         operatorType,
				operation: operationRightBrace,
			})
			i++
			continue
		}
		if s[i] == '+' {
			ret = append(ret, token{
				t:         operatorType,
				operation: operationAdd,
			})
			i++
			continue
		}
		if s[i] == '-' {
			ret = append(ret, token{
				t:         operatorType,
				operation: operationSub,
			})
			i++
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			j := i
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
			}
			v, err := strconv.ParseInt(s[i:j], 10, 64)
			if err != nil {
				return nil, err
			}
			ret = append(ret, token{
				t:     valueType,
				value: int(v),
			})
			i = j
			continue
		}
		i++
	}
	// fmt.Println(ret)
	return ret, nil
}
