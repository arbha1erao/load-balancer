package main

import (
	"github.com/arbha1erao/load-balancer/lb"
	"github.com/arbha1erao/load-balancer/utils"
)

type Config struct {
	Servers             []lb.Server `json:"servers"`
	HealthCheckInterval int         `json:"health_check_interval"`
	RetryCount          int         `json:"retry_count"`
	WebhookAlertURL     string      `json:"webhook_alert_url"`
}

func main() {
	utils.NewLogger()
	utils.Logger.Info().Msg("[MAIN] load balancer starting")

	var config Config

	if err := utils.LoadConfig("config.json", &config); err != nil {
		utils.Logger.Fatal().Err(err).Msg("[MAIN] failed to load configuration")
	}

	utils.Logger.Info().Msg("[MAIN] load balancer started")
}
