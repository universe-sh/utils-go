package types

import (
	"time"
)

// Timestamp int64
type Timestamp int64

// Int64toTimestamp converter
func Int64toTimestamp(t int64) Timestamp {
	return Timestamp(t)
}

// Int64 Timestamp converter
func (t Timestamp) Int64() int64 {
	return int64(t)
}

// Time Timestamp converter
func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}
