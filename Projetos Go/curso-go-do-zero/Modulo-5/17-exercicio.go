package main

// https://www.codewars.com/kata/56efc695740d30f963000557/train/go

// func main() {

// 	fmt.Println(ToAlternatingCase("Hello World"))
// }

// CODEWARS
// func ToAlternatingCase(str string) string {
// 	result := []rune{}
// 	for _, ch := range str {
// 		if unicode.IsUpper(ch) {
// 			result = append(result, unicode.ToLower(ch))
// 		} else if unicode.IsLower(ch) {
// 			result = append(result, unicode.ToUpper(ch))
// 		} else {
// 			result = append(result, ch)
// 		}
// 	}

// 	return string(result)
// }

// func ToAlternatingCase(str string) string {
// 	var result string

// 	for _, v := range str {
// 		if v >= 65 && v <= 90 {
// 			result += string(v + 32)
// 		} else if v >= 97 && v <= 122 {
// 			result += string(v - 32)
// 		} else {
// 			result += string(v)
// 		}
// 	}
// 	return result
// }

// ###################################

// https://www.codewars.com/kata/57f24e6a18e9fad8eb000296/train/go

// func main() {
// 	fmt.Println(HowMuchILoveYou(99841828))
// }

// func HowMuchILoveYou(i int) string {
// 	phrases := []string{
// 		"I love you",
// 		"a little",
// 		"a lot",
// 		"passionately",
// 		"madly",
// 		"not at all",
// 	}

// 	return phrases[(i-1)%6]
// }
