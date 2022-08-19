package main

import (
	"fmt"
)


type ColorType int

const (
	Red    ColorType = iota // 0
	Blue                    //1
	Green                   //2
	Yellow                  //3
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My Favorite Color is " + colorToString(getMyFavoriteColor()))

	// test fallthrough
	a := 3

	switch a {
	case 1:
		fmt.Println("a==1")
	case 2:
		fmt.Println("a==2")
	case 3:
		fmt.Println("a==3")
		fallthrough
	case 4:
		fmt.Println("a==4")
	case 5:
		fmt.Println("a==5")
	default:
		fmt.Println("Default: a > 5")
	}
}
