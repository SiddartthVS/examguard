package main

import (
	"fmt"
	"time"
)

func main() {
	event := Event{
		Time:    time.Now(),
		Source:  "USB",
		Type:    "USB_DEVICE_ADDED",
		Details: "A USB device was connected",
	}
	LogEvent(event)
	event = Event{
		Time:    time.Now(),
		Source:  "CtrlC/V",
		Type:    "COPIED/PASTED TEXT",
		Details: "Text was copied to clipboard",
	}
	LogEvent(event)

	hash := HashEvent(event.Details)
	fmt.Println("Hash of event details:", hash)

	err := VerifyChain()
	if err != nil {
		fmt.Println("Verification failed:", err)
	}
}
