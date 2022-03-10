package Helpers

import (
	"testing"
)

func TestComputeHmac256Signature(t *testing.T) {
	type args struct {
		message string
		secret  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test with sample data",
			args: args{
				message: "timestamp=1578963600000",
				secret:  "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
			},
			want: "d84e6641b1e328e7b418fff030caed655c266299c9355e36ce801ed14631eed4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeHmac256Signature(tt.args.message, tt.args.secret); got != tt.want {
				t.Errorf("ComputeHmac256Signature() = %v, want %v", got, tt.want)
			}
		})
	}
}
