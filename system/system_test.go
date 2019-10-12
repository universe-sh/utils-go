package system

import "testing"

func TestGetEnv(t *testing.T) {
	type args struct {
		name     string
		defValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{name: "ARG1", defValue: "test1"}, want: "test1"},
		{name: "test2", args: args{name: "ARG2", defValue: "test2"}, want: "test2"},
		{name: "test3", args: args{name: "ARG3", defValue: "test3"}, want: "test3"},
		{name: "test4", args: args{name: "ARG4", defValue: "test4"}, want: "test4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.args.name, tt.args.defValue); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
