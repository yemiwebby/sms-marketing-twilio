package sms

import (
	"fmt"
	"net/http"
	"strings"
)


func HandleIncomingSMS(w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("From")
	body := strings.TrimSpace(strings.ToLower(r.FormValue("Body")))

	customer := FindCustomer(from)
	if customer == nil {
		fmt.Printf("Customer with phone number %s not found\n", from)
		http.Error(w, "Customer not found", http.StatusNotFound)
		return 
	}

	if body == "stop" && !customer.OptedOut {
		customer.OptedOut = true
		fmt.Printf("Customer %s has opted out.\n", from)
	} else {
		fmt.Printf("Received message from %s: %s\n", from, body)
	}

	fmt.Fprintf(w, "<Response></Response>")
}