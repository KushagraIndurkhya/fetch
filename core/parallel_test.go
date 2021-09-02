package core

import "testing"

func Test_get_length(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_length(tt.args.url); got != tt.want {
				t.Errorf("get_length() = %v, want %v", got, tt.want)
			}
		})
	}
}
