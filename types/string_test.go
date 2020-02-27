package types

import "testing"

func TestMapArrayString(t *testing.T) {
	type args struct {
		a map[string][]string
		b string
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{a: map[string][]string{"test1": []string{"test0", "test1"}}, b: "test", c: 1}, want: ""},
		{args: args{a: map[string][]string{"test2": []string{"test0", "test1"}}, b: "test2", c: 1}, want: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapArrayString(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("MapArrayString() = %v, want %v", got, tt.want)
			}
		})
	}
}
