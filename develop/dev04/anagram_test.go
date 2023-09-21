package anagram

import (
	"reflect"
	"testing"
)

func Test_anagram(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{name: "test case",
			args: args{input: []string{
				"Пятак", "Пятак", "пятка", "тяпка",
				"листок", "слиток", "столик", "СТОЛИК",
				"без анаграммы"}},
			want: map[string][]string{
				"пятак":  {"пятка", "тяпка"},
				"листок": {"слиток", "столик"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagram(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
