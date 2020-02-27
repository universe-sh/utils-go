package env

import (
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{args: args{a: []string{}}, wantErr: false},
		{args: args{a: []string{"TEST"}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Check(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{a: "test1", b: "test2"}, want: "test2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
