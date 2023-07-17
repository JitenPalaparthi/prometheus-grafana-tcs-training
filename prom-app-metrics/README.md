# How to configure

- Make sure to install Golang
- Please follow https://go.dev/doc/install 
- git clone <this repo>
- cd prom-app-metrics


- The application by default works on 2110 port. To change the port you are free to go and change in source code. here in main.go

```
if err := http.ListenAndServe(":2110", nil); err != nil {
		log.Fatalln(err)
	}
```
- Build go binary 

```
go build -o prom-app main.go
```

- go to prometheus.yml file and add new job in scrap_configs sections as below

```
  - job_name: "prom-app"
    metrics_path: "/mymetrics"

    static_configs:
            - targets: ["localhost:2110"]
```
- Start prom-app

```
./prom-app

```

- To call ping and health request. These details are maintained as metrics.
- names of the metrics are as followed  ping: prom_app_metrics_ping_request_count
- health: prom_app_metrics_health_request_count

```
curl localhost:2110/ping
curl localhost:2110/health
```

