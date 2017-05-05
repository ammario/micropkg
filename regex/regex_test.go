package regex

import (
	"reflect"
	"regexp"
	"testing"
)

var r = regexp.MustCompile(`(?P<first>[0-9]+)\.(?P<second>[0-9]+)`)

func TestSubexp(t *testing.T) {
	type args struct {
		r      *regexp.Regexp
		target string
		subexp string
	}
	tests := []struct {
		name    string
		args    args
		wantVal string
	}{
		{"second", args{r, "5345345.911", "second"}, "911"},
		{"first", args{r, "5345345.911", "first"}, "5345345"},
		{"none", args{r, "5345345.911", "third"}, ""},
		{"empty", args{r, "", "first"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := Subexp(tt.args.r, tt.args.target, tt.args.subexp); gotVal != tt.wantVal {
				t.Errorf("Subexp() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestSubexpMap(t *testing.T) {
	type args struct {
		r      *regexp.Regexp
		target string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"basic", args{r, "16601.234234"}, map[string]string{
			"first":  "16601",
			"second": "234234",
		}},
		{"no second", args{r, "14"}, map[string]string{}},
		{"nothing", args{r, ""}, map[string]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubexpMap(tt.args.r, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubexpMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
