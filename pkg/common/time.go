// Copyright (C) Subhajit DasGupta 2021

package common

import (
	"fmt"
	"time"
)

// TimeFormat is the format used to express time as strings.
const TimeFormat = time.RFC3339Nano

// Time is a JSON and protobuf safe Time struct.
type Time struct {
	time.Time `protobuf:"-"`
}

// NewTime creates a new Time struct from the given time.Time.
func NewTime(t time.Time) Time {
	return Time{t}
}

// Now returns an analog of time.Now().
func Now() Time {
	t := Time{}
	t.Time = time.Now()

	return t
}

// UTC returns an analog of time.UTC().
func (t Time) UTC() Time {
	t.Time = t.Time.UTC()

	return t
}

// After returns true of the given Time v is after t.
func (t Time) After(v Time) bool {
	tt := t.Time
	vt := v.Time

	return tt.After(vt)
}

// Parse parses a string representation of time from val.
func (t *Time) Parse(val string) error {
	t1, err := time.Parse(TimeFormat, val)
	if err != nil {
		return err
	}

	t.setTimeNanoseconds(t1)

	return nil
}

// TextValue returns a string representation of t.
func (t Time) TextValue() string {
	if t.Time.IsZero() {
		return ""
	}

	return t.Time.Format(TimeFormat)
}

// MarshalJSON marshals t as a []byte.
func (t Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.TextValue())

	return []byte(stamp), nil
}

// UnmarshalJSON unmarshals the given []byte into t.
func (t *Time) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "\"\"" {
		t.setZeroTime()

		return nil
	}

	t1, err := time.Parse(`"`+TimeFormat+`"`, s)
	if err != nil {
		return err
	}

	t.setTimeNanoseconds(t1)

	return nil
}

// setTimeNanoseconds sets the given time to t.
func (t *Time) setTimeNanoseconds(t1 time.Time) {
	*t = Time{universalTimeNanosecond(t1)}
}

func (t *Time) setZeroTime() {
	*t = Time{time.Time{}}
}

// universalTimeNanosecond returns the universal time, rounded to nanoseconds,
// of the given time.
func universalTimeNanosecond(t time.Time) time.Time {
	return t.Round(time.Nanosecond).UTC()
}

// universalTimeMillisecond returns the universal time, rounded to milliseconds,
// of the given time.
func universalTimeMillisecond(t time.Time) time.Time {
	return t.Round(time.Millisecond).UTC()
}
