package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SetStatus(custom string) {
	discordToken := os.Getenv("DISCORD_TOKEN")
	status := "dnd" // or "idle", "dnd"

	payload, _ := json.Marshal(map[string]interface{}{
		"status": status,
		"custom_status": map[string]interface{}{ // Corrected field name
			"text": custom,
		},
	})
	req, _ := http.NewRequest("PATCH", "https://discord.com/api/v9/users/@me/settings", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", discordToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Discord status updated successfully!")
	} else {
		fmt.Println("Failed to update Discord status. Status code:", resp.StatusCode)
	}

}
