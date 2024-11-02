package main

import "fmt"

func main() {
	colors := map[string]string{
		"green": "#22fbs",
		"blue":  "#88215c",
	}
	fmt.Println(colors)
	segra := make(map[string]int)
	segra["one"] = 1
	delete(segra, "one")
	fmt.Println(segra)
	print(colors)
}
func print(c map[string]string){
	for color, hex := range c{
		fmt.Println(color, hex)
	}
}