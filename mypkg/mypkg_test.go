package mypkg

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Contains:", args{"This is a sentence", "is a"}, true},
		{"Not Contains:", args{"This is a sentence", "not"}, false},
		{"Empty str2:", args{"This is a sentence", ""}, true},
		{"Empty str1:", args{"", "is a"}, false},
		{"All empty:", args{"", ""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.args.str1, tt.args.str2)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
