package main

// import "fmt"

// func main() {
// 	fmt.Println(Points([]string{"1:0", "2:2", "3:3", "0:1", "4:0"}))
// }

// func Points(games []string) int {
// 	points := 0

// 	for _, v := range games {
// 		if v[0] > v[2] {
// 			points += 3
// 		}
// 		if v[0] == v[2] {
// 			points++
// 		}
// 	}
// 	return points
// }

//CODEWARS

// func Points(games []string) int {
//   result := 0
//   for _, game := range games {
//     str := strings.Split(game, ":")
//     x, y := str[0], str[1]
//     switch {
//       case x > y:
//         result += 3
//       case x == y:
//         result += 1
//     }
//   }
//   return result
// }
