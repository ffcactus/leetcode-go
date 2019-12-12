package integer_to_english_words

import "fmt"

var mapping = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func numberToWords(num int) string {
	s := fmt.Sprintf("%d", num)
	b := []byte(s)
	billion := ""
	if len(b) > 9 {
		part := b[:len(b)-9]
		fmt.Printf("billion part = %v\n", string(part))
		billion += convertPart(part)
		billion += " Billion"
		b = b[len(b)-9:]
	}

	million := ""
	if len(b) > 6 {
		part := b[:len(b)-6]
		fmt.Printf("million part = %v\n", string(part))
		t := convertPart(part)
		if t != "" {
			million += t + " Million"
		}
		b = b[len(b)-6:]
	}

	thousand := ""
	if len(b) > 3 {
		part := b[:len(b)-3]
		fmt.Printf("thousand part = %v\n", string(part))
		t := convertPart(part)
		if t != "" {
			thousand += t + " Thousand"
		}
		b = b[len(b)-3:]
	}
	fmt.Printf("rest part = %v\n", string(b))
	rest := convertPart(b)
	fmt.Printf("num = %d, billion = %v, million = %v, thousand = %v, rest = %v\n", num, billion, million, thousand, rest)
	out := ""
	if billion != "" {
		out += billion
	}
	if million != "" {
		if out != "" {
			out += " "
		}
		out += million
	}
	if thousand != "" {
		if out != "" {
			out += " "
		}
		out += thousand
	}
	if rest != "" {
		if out != "" {
			out += " "
		}
		out += rest
	}
	return out
}

func convertPart(part []byte) string {
	out := ""
	if len(part) == 3 {
		if part[0] != '0' {
			out += mapping[part[0]-'0']
			out += " Hundred"
			if suffix := convert00to99(string(part[1:])); suffix != "" {
				out += " " + suffix
			}
			return out
		}
		return convert00to99(string(part[1:]))
	}
	if len(part) == 2 {
		out += convert00to99(string(part))
		return out
	}
	if len(part) == 1 {
		out += mapping[part[0]-'0']
		return out
	}
	return out
}

func convert00to99(s string) string {
	out := ""
	switch s {
	case "00":
	case "01", "02", "03", "04", "05", "06", "07", "08", "09":
		out += mapping[s[1]-'0']
	case "10":
		out += "Ten"
	case "11":
		out += "Eleven"
	case "12":
		out += "Twelve"
	case "13":
		out += "Thirteen"
	case "14":
		out += "Fourteen"
	case "15":
		out += "Fifteen"
	case "16":
		out += "Sixteen"
	case "17":
		out += "Seventeen"
	case "18":
		out += "Eighteen"
	case "19":
		out += "Nineteen"
	case "20":
		out += "Twenty"
	case "21", "22", "23", "24", "25", "26", "27", "28", "29":
		out += "Twenty "
		out += mapping[s[1]-'0']
	case "30":
		out += "Thirty"
	case "31", "32", "33", "34", "35", "36", "37", "38", "39":
		out += "Thirty "
		out += mapping[s[1]-'0']
	case "40":
		out += "Forty"
	case "41", "42", "43", "44", "45", "46", "47", "48", "49":
		out += "Forty "
		out += mapping[s[1]-'0']
	case "50":
		out += "Fifty"
	case "51", "52", "53", "54", "55", "56", "57", "58", "59":
		out += "Fifty "
		out += mapping[s[1]-'0']
	case "60":
		out += "Sixty"
	case "61", "62", "63", "64", "65", "66", "67", "68", "69":
		out += "Sixty "
		out += mapping[s[1]-'0']
	case "70":
		out += "Seventy"
	case "71", "72", "73", "74", "75", "76", "77", "78", "79":
		out += "Seventy "
		out += mapping[s[1]-'0']
	case "80":
		out += "Eighty"
	case "81", "82", "83", "84", "85", "86", "87", "88", "89":
		out += "Eighty "
		out += mapping[s[1]-'0']
	case "90":
		out += "Ninety"
	case "91", "92", "93", "94", "95", "96", "97", "98", "99":
		out += "Ninety "
		out += mapping[s[1]-'0']
	}
	return out
}

func allZero(s []byte) bool {
	for _, v := range s {
		if v != '0' {
			return false
		}
	}
	return true
}
