package main

// import "fmt"

// func main() {

// 	fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
// }

// JEITO TOSCO QUE FIZ

// func CountPositivesSumNegatives(numbers []int) []int {

// 	var qtNumberPos int
// 	var sumnumberNeg int

// 	for _, number := range numbers {
// 		if number > 0 {
// 			qtNumberPos++
// 		}

// 		if number < 0 {
// 			sumnumberNeg += number
// 		}
// 	}

// 	res := []int{qtNumberPos, sumnumberNeg}

// 	return res
// }

// JEITO MAIS VOTADO DO CODEWARS
// func CountPositivesSumNegatives(numbers []int) []int {
//   res := []int{0,0}
//   for _, v := range numbers {
//     if v > 0 {
//       res[0] += 1
//     }else {
//       res[1] += v
//     }
//   }
//   return res // your code here
// }
