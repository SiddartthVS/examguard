package main

import (
	"encoding/json"
	"os"
)

var prevHash string

func LogEvent(event Event) error {

	record := LogRecord{
		Event:    event,
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
