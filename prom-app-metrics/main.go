package main

import (
	"io"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "prom_app_metrics_ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var healthCounter = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "prom_app_metrics_health_request_count",
		Help: "No of request handled by Health handler",
	},
)

var rootCounter = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "prom_app_metrics_root_request_count",
		Help: "No of request handled by root handler",
	},
)

func main() {

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootCounter.Inc()
		log.Println("root handler.Sending metrics to prometheus")
		io.WriteString(w, "Hello World")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingCounter.Inc()
		log.Println("ping handler.Sending metrics to prometheus")
		io.WriteString(w, "pong")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthCounter.Inc()
		log.Println("health handler.Sending metrics to prometheus")
		io.WriteString(w, "healthy")
	})

	log.Println("Server has started on port 2110")

	if err := http.ListenAndServe(":2110", nil); err != nil {
		log.Fatalln(err)
	}
}
