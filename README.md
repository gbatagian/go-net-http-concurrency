# Concurrency patterns with GO's net/http package

### Run Server
```bash
go run main.go
```

### Sample Sleep Request (0-1 Second Duration)
```bash
curl http://127.0.0.1:8080/sleep
```

### Call the /sleep endpoint concurrently n-times (last path parameter)
Simulates sending multiple requests to a third-party API. The semaphore design pattern is utilized to limit the number of concurrent requests running.
```bash
curl http://127.0.0.1:8080/sleep/100
```

### Description
This sample repository demonstrates concurrent patterns using Go's net/http package. The server listens on port 8080 and provides a `/sleep` endpoint that simulates a delay of 0-1 second.  The endpoint `/sleep/{n}` allows you to send concurrent requests to the `/sleep` endpoint, replacing {n} with the desired number of concurrent requests.
```bash 
⚡ curl http://127.0.0.1:8080/sleep/500
Benchmark sleep/500 time taken: 2.942936256s
Requests served: 500
Slowest requests time taken: 997ms
Requests total computation time: 3m52.872
```