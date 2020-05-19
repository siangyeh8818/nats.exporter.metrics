package main

import (
	"log"
	"os"

	. "github.com/pnetwork/sre.monitor.metrics/internal/app/openfaas"
)

func main() {
	log.Println("Prometheus exporter for nats-stream")
	nats_ip_env := os.Getenv("NATS_IP")
	if nats_ip_env == "" {
		nats_ip_env = "nats"
	}
	log.Printf("NATS_IP : %s \n", nats_ip_env)
	nats_port_env := os.Getenv("NATS_PORT")
	if nats_port_env == "" {
		nats_port_env = "8222"
	}
	log.Printf("NATS_PORT : %s \n", nats_port_env)
	Run_Exporter_Server()
}
