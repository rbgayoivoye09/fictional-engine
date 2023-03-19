package hashdemo

import "testing"

func Test_calculateHash(t *testing.T) {
	type args struct {
		toBeHash string
	}
	tests := []struct {
		name       string
		args       args
		wantHashed string
	}{
		{
			name: "test1",
			args: args{
				toBeHash: "welcome my golang chain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHashed := calculateHash(tt.args.toBeHash); gotHashed != tt.wantHashed {
				t.Errorf("calculateHash() = %v, want %v", gotHashed, tt.wantHashed)
			}
		})
	}
}
