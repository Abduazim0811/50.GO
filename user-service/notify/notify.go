package notify

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func NotifyOrderService(userID, action string) {
	url := "http://localhost:9001/order/notify"
	payload := map[string]string{"user_id": userID, "action": action}
	jsonPayload, _ := json.Marshal(payload)
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Failed to notify Order Service: %v\n", err)
		return
	}
	defer resp.Body.Close()
}
