# thanos-load-test

Env Variables:
* ENDPOINT (default = "prometheus.addon-metricset-ns.svc.cluster.local")
* PORT (default = "9090")
* GOROUTINES (default = 5)
* START_TIME (in unix epoch timestamp)
* END_TIME (in unix epoch timestamp)

Build Image:
```console
docker build --platform=linux/amd64 -t sample-metric .
```
