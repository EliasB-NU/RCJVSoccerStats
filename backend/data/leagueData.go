package data

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetLeagues(url string) *League {
	agent := fiber.Get(fmt.Sprintf("%sleagues", url)).InsecureSkipVerify()
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 || statusCode != 200 {
		log.Println(statusCode, errs)
		return nil
	}

	var leagues League
	err := json.Unmarshal(body, &leagues)
	if err != nil {
		log.Printf("Error unmarshalling leagues: %v\n", err)
		return nil
	}
	return &leagues
}
