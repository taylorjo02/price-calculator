package cmd

import "fmt"

type CMDManager struct {
	
}

func (cmd CMDManager) ReadLines() ([]string, error) {

	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		_, err := fmt.Scan(&price)
		if err != nil {
			fmt.Println(err)
		}




		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}