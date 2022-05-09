package nbminer

import "testing"

func Test_generalizeGpuModel(t *testing.T) {
	type args struct {
		nbDeviceInfo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "match with RTX Ti models",
			args: args{
				nbDeviceInfo: "NVIDIA GeForce RTX 3080 Ti LHR",
			},
			want: "3080 Ti",
		},
		{
			name: "match with RTX models",
			args: args{
				nbDeviceInfo: "NVIDIA GeForce RTX 3070 LHR",
			},
			want: "3070",
		},
		{
			name: "match with RTX non-LHR models",
			args: args{
				nbDeviceInfo: "NVIDIA GeForce RTX 3070",
			},
			want: "3070",
		},
		{
			name: "match with RTX non-LHR Ti models",
			args: args{
				nbDeviceInfo: "NVIDIA GeForce RTX 3070 Ti",
			},
			want: "3070 Ti",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cropGpuModel(tt.args.nbDeviceInfo); got != tt.want {
				t.Errorf("cropGpuModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
