package strutil

import "testing"

func TestEllipsis(t *testing.T) {
	type args struct {
		str    string
		maxLen int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fast path", args{"poop", 10}, "poop"},
		{"exactly", args{"0123456789", 10}, "0123456789"},
		{"over", args{"0123456789a", 10}, "0123456..."},
		{"utf8 fast", args{"Прив", 10}, "Прив"},
		{"utf8 exactly", args{"ПривПривив", 10}, "ПривПривив"},
		{"utf8 over", args{"ПривПрививПривПривив", 10}, "ПривПри..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ellipsis(tt.args.str, tt.args.maxLen); got != tt.want {
				t.Errorf("Ellipsis() = %v, want %v", got, tt.want)
			}
		})
	}
}
