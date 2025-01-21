package main

// import "fmt"

// func main() {
// 	fmt.Println(CalculateYears(5))
// }

// func CalculateYears(years int) (result [3]int) {
// 	if years == 1 {
// 		return [3]int{1, 15, 15}
// 	}
// 	if years == 2 {
// 		return [3]int{2, 24, 24}
// 	}

// 	catYears := 24 + (years-2)*4
// 	dogYears := 24 + (years-2)*5

// 	return [3]int{years, catYears, dogYears}

// }

//CODEWARS
// func CalculateYears(years int) (result [3]int) {
//   switch years {
//     case 1: result = [3]int{1, 15, 15}
//     case 2: result = [3]int{2, 24, 24}
//     default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}
//   }
//   return
// }
