package logging

import (
	"encoding/json"
	"os"
	"examguard/internal/event"
)

var prevHash string

func LogEvent(ev event.Event) error {

	record := event.LogRecord{
		ComputerID: "Computer-1",
		Event:    ev,
		PrevHash: prevHash,
	}

	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	record.Hash = HashEvent(string(data))

	data, err = json.Marshal(record)
	if err != nil {
		return err
	}
	data = append(data, '\n')

	file, err := os.OpenFile("events_chain.jsonl", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	prevHash = record.Hash
	return nil
}
