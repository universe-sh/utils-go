package slice

import (
	"testing"
)

func TestStringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{"test1", []string{"test1", "test2", "test3"}}, want: true},
		{args: args{"t%^&", []string{"test2", "t%^&%", "test3"}}, want: false},
		{args: args{"1234", []string{"test2", "test3", "1234"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapCounterInSlice(t *testing.T) {
	type args struct {
		a    map[string]int64
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{map[string]int64{"test2": 2, "test1": 1}, []string{"test2", "test2", "test3"}}, want: false},
		{args: args{map[string]int64{"test24": 1, "test3": 1}, []string{"test24", "test2", "test3"}}, want: true},
		{args: args{map[string]int64{"test24": 1, "test3": 1}, []string{"test24"}}, want: false},
		{args: args{map[string]int64{}, []string{"test24", "test2", "test3"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapCounterInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("MapCounterInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
