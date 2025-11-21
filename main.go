package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin) 


	for {
	// Prompt user for mobile money number
	fmt.Print("Enter your mobile money number: ")
	mobileNumber, _ := reader.ReadString('\n')
	mobileNumber = strings.TrimSpace(mobileNumber) // remove newline
	fmt.Print(mobileNumber)

	if !isCameroonNumber(mobileNumber) {
		fmt.Println("Invalid Cameroon mobile number. It should start with '237' and be 12 digits long.")
		continue
	}

	// Prompt user for amount
	fmt.Print("Enter amount: ")
	amount, _ := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)

	// Prompt user for description
	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	//reader.ReadString('\n') // Wait for user to press Enter

	
	fmt.Println("\n--- Transaction Summary ---")
	fmt.Println("Mobile Money Number:", mobileNumber)
	fmt.Println("Amount:", amount)
	fmt.Println("Description:", description)

	
	resp, err := MakeMTNMobileTransferFromCampayAccount(CollectRequest{
		PhoneNumber: mobileNumber,
		Amount:      amount,
		Description: description,
	})

	if err != nil {
		fmt.Println("Error:", err)
		continue
	}

	if resp.Status != "success" {
		fmt.Println("Transaction failed:", resp.Message)
		continue
	} else {
		fmt.Println("Transaction completed successfully", resp)
	}


	//fmt.Println("Status: Transaction completed successfully", resp)

	// Ask if user wants to do another transaction
		fmt.Print("Do you want to make another transaction? (y/n): ")
		again, _ := reader.ReadString('\n')
		again = strings.TrimSpace(strings.ToLower(again))
		if again != "y" {
			fmt.Println("Exiting. Thank you!")
			break
		}

		fmt.Println() 
	}



}