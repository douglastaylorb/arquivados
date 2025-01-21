package main

import (
	"calculator/math"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	fmt.Println(e)

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.Sub(80, 1)
	fmt.Println(sub)
}
