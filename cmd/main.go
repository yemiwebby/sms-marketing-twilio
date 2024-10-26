package main

import (
	"fmt"
	"log"
	"net/http"
	"sms-marketing-with-sdk/config"
	"sms-marketing-with-sdk/internal/campaign"
	"sms-marketing-with-sdk/internal/optout"
)

func main() {
	config.LoadEnv()

	http.HandleFunc("/send-campaign", sendCampaignHandler)
	http.HandleFunc("/process-reply", optout.HandleIncomingSMS)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func sendCampaignHandler(w http.ResponseWriter, r *http.Request) {
	message := "Exclusive Offer: Get 20% off your next purchase!"
	if err := campaign.SendMarketingCampaign(message); err != nil {
		http.Error(w, "Failed to send campaign", http.StatusInternalServerError)
		return
	}
	
	fmt.Fprintln(w, "Campaign sent successfully")
}