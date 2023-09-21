package unpack

import "testing"

func TestDoUnpack(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				input: "a4bc2d5e",
			},
			want: "aaaabccddddde",
		},
		{
			name: "test2",
			args: args{
				input: "a11bc2d5e",
			},
			want: "aaaaaaaaaaabccddddde",
		},
		{
			name: "test3",
			args: args{
				input: "a4bc2d5e",
			},
			want: "aaaabccddddde",
		},
		{
			name: "test4",
			args: args{
				input: "abcd",
			},
			want: "abcd",
		},
		{
			name: "test5",
			args: args{
				input: "45",
			},
			want: "",
		},
		{
			name: "test6",
			args: args{
				input: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoUnpack(tt.args.input); got != tt.want {
				t.Errorf("DoUnpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
