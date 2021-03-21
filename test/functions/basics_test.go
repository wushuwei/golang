package functions

import (
	"testing"
)

func TestPointers(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pointers()
		})
	}
}

func TestFunctions(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{0, 0},
			0,
			0,
			false,
		},
		{
			"test2",
			args{0, 0},
			0,
			0,
			true,
		},
		{
			"test3",
			args{1, 0},
			0,
			0,
			false,
		},
		{
			"test4",
			args{1, 0},
			1,
			1,
			false,
		},
		{
			"test5",
			args{1, 1},
			1,
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Functions(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("Functions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Functions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Functions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
