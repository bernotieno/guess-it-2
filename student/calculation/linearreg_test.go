package calculation

import "testing"

func Test_linearRegression(t *testing.T) {
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
			want:  1, 
			want1: 1, 
		},
		{
			name: "Negative slope",
			args: args{
				data: []float64{5, 4, 3, 2, 1},
			},
			want:  -1, // slope
			want1: 5,  // intercept
		},
		{
			name: "Two points",
			args: args{
				data: []float64{2, 4},
			},
			want:  2, // slope
			want1: 2, // intercept
		},
		{
			name: "Zero variance",
			args: args{
				data: []float64{3, 3, 3, 3, 3},
			},
			want:  0, // slope
			want1: 3, // intercept (constant value)
		},
		{
			name: "General case",
			args: args{
				data: []float64{1, 3, 2, 4, 5},
			},
			want:  0.9, // approximate slope
			want1: 1.2,   // approximate intercept
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := linearRegression(tt.args.data)
			if got != tt.want {
				t.Errorf("linearRegression() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("linearRegression() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
