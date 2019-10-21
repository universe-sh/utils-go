package math

import (
	"testing"
)

func TestPourcent(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{a: 1, b: 1}, want: 100},
		{args: args{a: -1, b: 2}, want: -50},
		{args: args{a: 100, b: 259}, want: 38.61003861003861},
		{args: args{a: 50, b: 1000000}, want: 0.005},
		{args: args{a: 385, b: 400}, want: 96.25},
		{args: args{a: 858, b: 1940}, want: 44.22680412371134},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pourcent(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Pourcent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{args: args{a: 1, b: -1}, want: 2},
		{args: args{a: -1, b: -32}, want: 31},
		{args: args{a: 100456, b: 259}, want: 100197},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
