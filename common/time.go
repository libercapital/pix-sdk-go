package common

import (
	"fmt"
	"strings"
	"time"
)

type PixTime struct {
	Time time.Time
}

var formatLayout = "2006-01-02T15:04:05"

func (p *PixTime) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(time.RFC3339, strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}
	p.Time = parsedTime.Truncate(time.Second).UTC()
	return nil
}

func (p PixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", p.Time.Format(formatLayout))), nil
}

func (p *PixTime) SetTime(newTime time.Time) {
	p.Time = newTime
}
