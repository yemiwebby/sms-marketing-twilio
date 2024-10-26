package campaign

import (
	"fmt"
	"log"
	"sms-marketing-with-sdk/internal/sms"
)

type Customer struct {
	PhoneNumber string
	OptedOut bool
}

var Customers = []*Customer{
	{PhoneNumber: "+447776736645", OptedOut: false},
	{PhoneNumber: "+0987654321", OptedOut: true},
}

func FindCustomer(phone string) *Customer {
	fmt.Println(phone, "PASSED NUMBER")
	for _, customer := range Customers {
		fmt.Println(customer.PhoneNumber, "PHONE NUMBER")
		if customer.PhoneNumber == phone {
			return customer
		}
	}
	return nil
}

func SendMarketingCampaign(message string) error {
	for _, customer := range Customers {
		if !customer.OptedOut {
			err := sms.SendSMS(customer.PhoneNumber, message)
			if err != nil {
				log.Printf("Failed to send SMS to %s: %v\n", customer.PhoneNumber, err)
				return fmt.Errorf("failed to send campaign: %w", err)
			}
		} else {
			fmt.Printf("Customer %s has opted out. Skipping.\n", customer.PhoneNumber)
		}
	}
	return nil
}