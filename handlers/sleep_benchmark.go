package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func makeRequest(c chan time.Duration, sem *chan bool) {
	defer func() { <-*sem }()

	rsp, err := http.Get("http://127.0.0.1:8080/sleep")
	if err != nil {
		fmt.Println(err.Error())
	}

	rspBody := ResponseSchema{}
	defer rsp.Body.Close()
	json.NewDecoder(rsp.Body).Decode(&rspBody)

	dur, _ := time.ParseDuration(rspBody.Time)
	c <- dur
}

func SleepN(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	var rqsTime time.Duration
	var m time.Duration
	c := make(chan time.Duration)
	sem := make(chan bool, 100)

	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid value `%s` provided as int", r.PathValue("n")), http.StatusBadRequest)
		return
	}

	receiveCnt := 0
	sendCnt := 0
	for receiveCnt < n {
		select {
		case sem <- true:
			if sendCnt < n {
				sendCnt++
				go makeRequest(c, &sem)
			}
		case dur := <-c:
			receiveCnt++
			rqsTime += dur
			if dur > m {
				m = dur
			}
		}
	}
	log.Println(receiveCnt, sendCnt)

	w.Write([]byte(fmt.Sprintf("Benchmark sleep/%d time taken: %v", n, time.Since(t))))
	w.Write([]byte(fmt.Sprintf("\nRequests served: %d", receiveCnt)))
	w.Write([]byte(fmt.Sprintf("\nSlowest requests time taken: %v", m)))
	w.Write([]byte(fmt.Sprintf("\nRequests total computation time: %v", rqsTime)))
}
