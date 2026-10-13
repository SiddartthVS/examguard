package main

import "time"

type Event struct {
	Time    time.Time
	Source  string
	Type    string
	Details string
}

type LogRecord struct {
	Event    Event  `json:"event"`
	PrevHash string `json:"prev_hash"`
	Hash     string `json:"hash"`
}
