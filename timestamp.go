package main

import (
	"encoding/json"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type Timestamp time.Time

func (timestamp *Timestamp) UnmarshalJSON(data []byte) error {
	var representation string

	err := json.Unmarshal(data, &representation)
	if err != nil {
		return err
	}

	value, err := time.Parse(TimeFormat, representation)
	if err != nil {
		return err
	}

	*timestamp = Timestamp(value)

	return nil
}
