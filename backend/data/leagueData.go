package data

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetLeagues(url string) *League {
	response, err := catigoalClient.Get(fmt.Sprintf("%sleagues", url))
	if err != nil {
		log.Printf("Error fetching Catigoal data: %v\n", err)
		return nil
	}
	defer response.Close()
	statusCode, body := response.StatusCode(), response.Body()
	if statusCode != 200 {
		log.Printf("Error getting leagues with code %d\n", statusCode)
		return nil
	}

	var leagues League
	err = json.Unmarshal(body, &leagues)
	if err != nil {
		log.Printf("Error unmarshalling leagues: %v\n", err)
		return nil
	}
	return &leagues
}
