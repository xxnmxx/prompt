package main

import "fmt"

const (
	black       = "\x1b[30m%v\x1b[0m"
	red         = "\x1b[31m%v\x1b[0m"
	green       = "\x1b[32m%v\x1b[0m"
	yellow      = "\x1b[33m%v\x1b[0m"
	blue        = "\x1b[34m%v\x1b[0m"
	magenta     = "\x1b[35m%v\x1b[0m"
	cyan        = "\x1b[36m%v\x1b[0m"
	white       = "\x1b[37m%v\x1b[0m"
	brightgreen = "\x1b[92m%v\x1b[0m"
)

func main() {
	hello := "hello colors!"
	fmt.Printf(black, hello)
	fmt.Println("")
	fmt.Printf(red, hello)
	fmt.Println("")
	fmt.Printf(green, hello)
	fmt.Println("")
	fmt.Printf(yellow, hello)
	fmt.Println("")
	fmt.Printf(blue, hello)
	fmt.Println("")
	fmt.Printf(magenta, hello)
	fmt.Println("")
	fmt.Printf(cyan, hello)
	fmt.Println("")
	fmt.Printf(white, hello)
	fmt.Println("")
	fmt.Printf(brightgreen, hello)
	fmt.Println("\x1b[92mHello Row Strings\x1b[0m")
}
