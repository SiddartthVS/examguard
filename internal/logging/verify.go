package logging

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"examguard/internal/event"
)

func VerifyChain() error {
	file, err := os.Open("events_chain.jsonl")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	prevHash := ""
	count := 0

	for reader.Scan() {
		var record event.LogRecord

		err := json.Unmarshal(reader.Bytes(), &record)
		if err != nil {
			return err
		}

		if record.PrevHash != prevHash {
			return fmt.Errorf("chain broken at record %d", count+1)
		}

		oldHash := record.Hash
		record.Hash = ""

		data, err := json.Marshal(record)
		if err != nil {
			return err
		}

		newHash := HashEvent(string(data))

		if oldHash != newHash {
			return fmt.Errorf("tampering detected at record %d", count+1)
		}

		prevHash = oldHash
		count++
	}

	if err := reader.Err(); err != nil {
		return err
	}

	fmt.Println("Chain verified successfully!")
	fmt.Println("Records verified:", count)

	return nil
}
