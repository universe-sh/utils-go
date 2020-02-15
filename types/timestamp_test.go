package types

import (
	"reflect"
	"testing"
	"time"
)

func TestInt64toTimestamp(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want Timestamp
	}{
		{args: args{t: 1}, want: Timestamp(1)},
		{args: args{t: 10000}, want: Timestamp(10000)},
		{args: args{t: 99999999}, want: Timestamp(99999999)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64toTimestamp(tt.args.t); got != tt.want {
				t.Errorf("Int64toTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestamp_Int64(t *testing.T) {
	tests := []struct {
		name string
		t    Timestamp
		want int64
	}{
		{t: Timestamp(1), want: 1},
		{t: Timestamp(-1), want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Int64(); got != tt.want {
				t.Errorf("Timestamp.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestamp_Time(t *testing.T) {
	tests := []struct {
		name string
		t    Timestamp
		want time.Time
	}{
		{t: Timestamp(1), want: time.Unix(1, 0)},
		{t: Timestamp(1000), want: time.Unix(1000, 0)},
		{t: Timestamp(99999999), want: time.Unix(99999999, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timestamp.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
