package models

import (
	"encoding/json"
	"time"
)

//layaout RFC3339 with -0700 (Numeric Time Zone)
const ctLayout = "2006-01-02T15:04:05-0700"

// Time: Custom Type
type Time struct {
	time.Time
}

// Time implements the Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(ctLayout))
}

// Time implements the Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}
	t.Time, err = time.Parse(ctLayout, s)
	t.Time = t.Time.UTC()
	return err
}
