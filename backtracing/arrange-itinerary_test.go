package backtracking

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	type args struct {
		startSite string
		tickets   [][2]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "5 site and 4 tickets",
			args: args{
				startSite: "JFK",
				tickets: [][2]string{
					{"MUC", "LHR"},
					{"JFK", "MUC"},
					{"SFO", "SJC"},
					{"LHR", "SFO"},
				},
			},
			want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name: "6 site and 7 tickets",
			args: args{
				startSite: "JFK",
				tickets: [][2]string{
					{"JFK", "SFO"},
					{"JFK", "ATL"},
					{"SFO", "ATL"},
					{"ATL", "JFK"},
					{"ATL", "SFO"},
				},
			},
			want: []string{"JFK", "SFO", "ATL", "JFK", "ATL", "SFO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindItinerary(tt.args.startSite, tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
