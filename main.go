package main

import (
	"log"
	"os"
	
	"examguard/internal/logging"
	"examguard/internal/monitor/usb"
)

func main() {
	log.Println("ExamGuard starting")

	if _, err := os.Stat("events_chain.jsonl"); err == nil {
		if err := logging.VerifyChain(); err != nil {
			log.Printf("Log verification failed: %v", err)
		} else {
			log.Println("Log integrity verified")
		}
	}

	go usb.MonitorUSB()

	log.Println("USB monitoring started")

	select {}
}