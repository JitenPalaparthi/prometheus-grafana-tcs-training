package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()

	pingClicked := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_ping_handler_click_count",
		Help: "The total number of times click pressed",
	})

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingClicked.Inc()
		fmt.Fprintln(w, "pong")

	})
	fmt.Println("Server started on port 2112")
	http.ListenAndServe(":2112", nil)
}
