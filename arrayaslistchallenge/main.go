package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func options() {
	fmt.Println("Available actions:")
	fmt.Println("0 - to shutdown")
	fmt.Println("1 - to add item(s) to list (comma delimited list)")
	fmt.Println("2 - to remove any items (comma delimited list)")
	fmt.Println("Enter a number for which action you want to do")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	groceryList := []string{}

	for {
		options()
		fmt.Print("Enter what you want to do: ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			fmt.Print("Enter item(s) to add (comma delimited): ")
			scanner.Scan()
			itemsToAdd := strings.Split(scanner.Text(), ",")

			for _, item := range itemsToAdd {
				item = strings.TrimSpace(item)
				if item != "" && !contains(groceryList, item) {
					groceryList = append(groceryList, item)
				} else if item != "" {
					fmt.Printf("Item \"%s\" is already in the list or is invalid.\n", item)
				}
			}
			sort.Strings(groceryList)
			fmt.Println("Current list:", groceryList)

		case "2":
			fmt.Print("Enter item(s) to remove (comma delimited): ")
			scanner.Scan()
			itemsToRemove := strings.Split(scanner.Text(), ",")

			for _, item := range itemsToRemove {
				item = strings.TrimSpace(item)
				index := indexOf(groceryList, item)
				if index != -1 {
					groceryList = append(groceryList[:index], groceryList[index+1:]...)
				} else {
					fmt.Printf("Item \"%s\" not found in the list.\n", item)
				}
			}
			sort.Strings(groceryList)
			fmt.Println("Current list:", groceryList)

		case "0":
			fmt.Println("Shutting down...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}
