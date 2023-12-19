package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill(reader *bufio.Reader) bill {
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill, reader *bufio.Reader) {
	option, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		parsed, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b, reader)
		}
		b.addItem(name, parsed)
		fmt.Println("Item added -", name, price)

		promptOptions(b, reader)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		parsed, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b, reader)
		}
		b.updateTip(parsed)
		fmt.Println("Tip added -", parsed)

		promptOptions(b, reader)
	case "s":
		b.save()
	default:
		fmt.Println("invalid option selected...")
		promptOptions(b, reader)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	myBill := createBill(reader)

	promptOptions(myBill, reader)
}
