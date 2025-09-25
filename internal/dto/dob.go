package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type DOB time.Time

const dobLayout = "2006-01-02" // YYYY-MM-DD

// Unmarshal JSON string ("YYYY-MM-DD") → time.Time
func (d *DOB) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s == "" {
		// allow empty DOB
		*d = DOB(time.Time{})
		return nil
	}

	t, err := time.Parse(dobLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}

	*d = DOB(t)
	return nil
}

// Marshal time.Time → JSON string ("YYYY-MM-DD")
func (d DOB) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`""`), nil // empty string
	}
	// wrap manually in quotes
	return []byte(t.Format(dobLayout)), nil
}


// Helper to cast to time.Time
func (d DOB) ToTime() time.Time {
	return time.Time(d)
}
