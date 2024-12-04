package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
/*Validates the request of the receipt
if any part of the validation fails the request will fail
*/
func validate_request(receipt Receipt) bool {
	if !validate_retailer(receipt.Retailer) {
		fmt.Println("Failure to validate Retailer")
		return false
	}
	if !validate_purchaseDate(receipt.PurchaseDate) {

		fmt.Println("Failure to validate purchase date")
		return false
	}
	if !validate_purchaseTime(receipt.PurchaseTime) {
		fmt.Println("Failure to validate purchase time")
		return false
	}
	if !validate_total(receipt.Total) {
		fmt.Println("Failure to validate total")
		return false
	}
	if !validate_items(receipt.Items) {
		fmt.Println("Failure to validate items")
		return false
	}
	return true

}

/*
retailer:

	description: The name of the retailer or store the receipt is from.
	type: string
	pattern: "^[\\w\\s\\-&]+$"
	example: "M&M Corner Market"
*/
//Validates the retailer if the string doesnt match the pattern or if the retailer is emtpy
func validate_retailer(retailer string) bool {
	if retailer == "" {
		return false
	}
	
	m, err := regexp.MatchString("^[\\w\\s\\-&]+$", retailer)
	if err != nil {
		fmt.Println("faulty regex")
		return false
	}
	if m {
		return true
	} else {
		return false
	}
}

/*
purchaseDate:

	description: The date of the purchase printed on the receipt.
	type: string
	format: date
	example: "2022-01-01"
*/
//must fit the format of the date otherwise returns false
func validate_purchaseDate(purchaseDate string) bool {
	if purchaseDate == "" {
		return false
	}
	var purchaseDateSplit = strings.Split(purchaseDate, "-")
	//if the split string doesnt equal 3 then it must be missing something so fail
	if len(purchaseDateSplit) != 3 {
		return false
	}
	//if the string contains anything other than numbers after dashes are removed then fail
	if !regexp.MustCompile(`\d`).MatchString(strings.Join(purchaseDateSplit, "")) {
		return false
	}
	day, hourErr := strconv.ParseInt(purchaseDateSplit[2], 10, 16)
	month, minuteErr := strconv.ParseInt(purchaseDateSplit[1], 10, 16)
	if hourErr != nil || minuteErr != nil {
		return false
	}
	//days cant be more than 31 or less than 0 and months cant be more than 12
	if day > 31 || day < 0 || month < 0 || month > 12 || len(purchaseDateSplit[0]) != 4 || len(purchaseDateSplit[1]) != 2 || len(purchaseDateSplit[2]) != 2 {
		return false
	}

	return true
}

/*
purchaseTime:

	description: The time of the purchase printed on the receipt. 24-hour time expected.
	type: string
	format: time
	example: "13:01"
*/

func validate_purchaseTime(purchaseTime string) bool {
	if purchaseTime == "" {
		return false
	}
	var purchaseTimeSplit = strings.Split(purchaseTime, ":")
	if len(purchaseTimeSplit) != 2 {
		return false
	}

	if !regexp.MustCompile(`\d`).MatchString(strings.Join(purchaseTimeSplit, "")) {
		return false
	}

	hour, hourErr := strconv.ParseInt(purchaseTimeSplit[0], 10, 16)
	minute, minuteErr := strconv.ParseInt(purchaseTimeSplit[1], 10, 16)
	if hourErr != nil || minuteErr != nil {
		return false
	}

	if hour > 24 || hour < 0 || minute > 60 || minute < 0 || len(purchaseTimeSplit[0]) != 2 || len(purchaseTimeSplit[1]) != 2 {
		return false
	}

	return true

}

/*
total:

	description: The total amount paid on the receipt.
	type: string
	pattern: "^\\d+\\.\\d{2}$"
	example: "6.49"
*/
func validate_total(total string) bool {
	if !regexp.MustCompile("^\\d+\\.\\d{2}$").MatchString(total) && total != "" {
		return false
	}
	return true
}

/*
items:
	type: array
	minItems: 1
	items:
		$ref: "#/components/schemas/Item"
*/
/*
Item:
	type: object
	required:
		- shortDescription
		- price
	properties:
		shortDescription:
			description: The Short Product Description for the item.
			type: string
			pattern: "^[\\w\\s\\-]+$"
			example: "Mountain Dew 12PK"
		price:
			description: The total price payed for this item.
			type: string
			pattern: "^\\d+\\.\\d{2}$"
			example: "6.49"

*/
func validate_items(items []Item) bool {
	if len(items) == 0 {
		return false
	}

	for _, a := range items {
		//loop through items and check to see if its null or if the string doesnt match the patterns 
		if a.Price == "" || a.ShortDescription == "" {
			return false
		}
		
		if !regexp.MustCompile("^\\d+\\.\\d{2}$").MatchString(a.Price) || !regexp.MustCompile("^[\\w\\s\\-]+$").MatchString(a.ShortDescription) {
			return false
		}
	}

	return true
}
