package calculation

import "testing"

func Test_calculateMean(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive numbers",
			args: args{
				data: []float64{1, 2, 3, 4, 5},
			},
			want: 3, 
		},
		{
			name: "Negative numbers",
			args: args{
				data: []float64{-1, -2, -3, -4, -5},
			},
			want: -3, 
		},
		{
			name: "Mixed numbers",
			args: args{
				data: []float64{-1, 0, 1},
			},
			want: 0, 
		},
		{
			name: "Single value",
			args: args{
				data: []float64{42},
			},
			want: 42, 
		},
		{
			name: "Large numbers",
			args: args{
				data: []float64{1000000, 2000000, 3000000},
			},
			want: 2000000, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMean(tt.args.data); got != tt.want {
				t.Errorf("calculateMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
