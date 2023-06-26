package demo34

import (
	"fmt"
)

type Color string

func main() {
	colorMap := make(map[Color]string)
	colorMap[Color("red")] = "#FF0000"
	colorMap[Color("green")] = "#00FF00"
	colorMap[Color("blue")] = "#0000FF"

	fmt.Println(colorMap[Color("red")])   // 输出: #FF0000
	fmt.Println(colorMap[Color("green")]) // 输出: #00FF00
	fmt.Println(colorMap[Color("blue")])  // 输出: #0000FF
}
