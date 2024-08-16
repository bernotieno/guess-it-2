package calculation

import (
	"testing"
)

func TestCalculateRange(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "Basic linear data",
			args: args{
				data: []float64{1, 2, 3, 4, 5},
			},
			want:  2.3569858633269067,
			want1: 9.643014136673093,
		},
		{
			name: "Decreasing series",
			args: args{
				data: []float64{5, 4, 3, 2, 1},
			},
			want:  -3.6430141366730933,
			want1: 3.6430141366730933,
		},
		{
			name: "Two points",
			args: args{
				data: []float64{1, 2},
			},
			want:  1.712, // Expected lower bound (approximated for the test case)
			want1: 4.288, // Expected upper bound (approximated for the test case)
		},
		{
			name: "Zero variance",
			args: args{
				data: []float64{3, 3, 3, 3, 3},
			},
			want:  3, // Lower bound (constant value)
			want1: 3, // Upper bound (constant value)
		},
		{
			name: "General case",
			args: args{
				data: []float64{1, 2, 3, 4, 5, 6},
			},
			want:  2.6006424711480127, // Expected lower bound (approximated for the test case)
			want1: 11.399357528851986, // Expected upper bound (approximated for the test case)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CalculateRange(tt.args.data)
			if got != tt.want {
				t.Errorf("CalculateRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalculateRange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
