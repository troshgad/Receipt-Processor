package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

	fmt.Println("receipt is valid")
	return true

}

/*
retailer:

	description: The name of the retailer or store the receipt is from.
	type: string
	pattern: "^[\\w\\s\\-&]+$"
	example: "M&M Corner Market"
*/
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
func validate_purchaseDate(purchaseDate string) bool {
	if purchaseDate == "" {
		fmt.Println("Hit 0")
		return false
	}
	var purchaseDateSplit = strings.Split(purchaseDate, "-")
	if len(purchaseDateSplit) != 3 {
		fmt.Println("Hit 1")
		return false
	}

	if !regexp.MustCompile(`\d`).MatchString(strings.Join(purchaseDateSplit, "")) {
		fmt.Println("Hit 2")
		return false
	}
	day, hourErr := strconv.ParseInt(purchaseDateSplit[2], 10, 16)
	month, minuteErr := strconv.ParseInt(purchaseDateSplit[1], 10, 16)
	if hourErr != nil || minuteErr != nil {
		fmt.Println("Hit 3")
		return false
	}

	if day > 31 || day < 0 || month < 0 || month > 12 || len(purchaseDateSplit[0]) != 4 || len(purchaseDateSplit[1]) != 2 || len(purchaseDateSplit[2]) != 2 {
		fmt.Println(day)
		fmt.Println(month)
		fmt.Println("Hit 4")
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
		if a.Price == "" || a.ShortDescription == "" {
			return false
		}

		if !regexp.MustCompile("^\\d+\\.\\d{2}$").MatchString(a.Price) || !regexp.MustCompile("^[\\w\\s\\-]+$").MatchString(a.ShortDescription) {
			return false
		}
	}

	return true
}
