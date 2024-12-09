package main

import (
	"log"
	"net/http"

	"github.com/arbha1erao/load-balancer/health"
	"github.com/arbha1erao/load-balancer/lb"
	"github.com/arbha1erao/load-balancer/routing"
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

	loadBalancer := &lb.LoadBalancer{Servers: make([]*lb.Server, len(config.Servers))}
	for i, s := range config.Servers {
		loadBalancer.Servers[i] = &lb.Server{URL: s.URL, Weight: s.Weight, Active: true}
	}

	go health.HealthCheck(loadBalancer.Servers, config.HealthCheckInterval, config.WebhookAlertURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := routing.ForwardRequest(loadBalancer, r, config.RetryCount)
		if err != nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
	})

	utils.Logger.Info().Msg("[MAIN] load balancer started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
