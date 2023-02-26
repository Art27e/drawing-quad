package quad

import (
	"fmt"
	"strings"
)

var mapColor = map[string]string{
	"red":    "\u001b[0;91m",
	"green":  "\u001b[32m",
	"yellow": "\u001b[33m",
	"blue":   "\u001b[34m",
	"orange": "\u001b[31m",
	"cyan":   "\u001B[36m",
	"purple": "\u001B[35m",
	"white":  "\u001B[37m",
	"black":  "\u001B[30m",
	"reset":  "\u001b[0m",
}

func Drawing(x, y int, symbolxC string, symbolxL string, symbolY string, color string) {
	if x <= 0 || y <= 0 {
		return
	}
	var mySlice1 []string
	mySlice1 = append(mySlice1, symbolxC)
	for i := 0; i < x-2; i++ {
		mySlice1 = append(mySlice1, symbolxL)
	}
	if x != 1 {
		mySlice1 = append(mySlice1, symbolxC)
	}
	justString := strings.Join(mySlice1, "")
	fmt.Println(string(mapColor[color]),justString, string(mapColor["reset"]) )
	// --------------------- //
	var mySlice2 []string // |  |
	mySlice2 = append(mySlice2, symbolY)
	for i := 0; i < x-2; i++ {
		mySlice2 = append(mySlice2, " ")
	}
	if x != 1 {
		mySlice2 = append(mySlice2, symbolY)
	}
	justString2 := strings.Join(mySlice2, "")
	// Vertical //
	i := 1
	for i < y-1 { // убираем 1 о
		i++
		fmt.Println(string(mapColor[color]),justString2,string(mapColor["reset"]))
	}
	// -------------------- //
	var mySlice3 []string
	mySlice3 = append(mySlice3, symbolxC)
	for i := 0; i < x-2; i++ {
		mySlice3 = append(mySlice3, symbolxL)
	}
	if x != 1 {
		mySlice3 = append(mySlice3, symbolxC)
	}
	justString3 := strings.Join(mySlice3, "")
	if y != 1 {
		fmt.Println(string(mapColor[color]),justString3,string(mapColor["reset"]))
	}
}
