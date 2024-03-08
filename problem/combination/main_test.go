package main

import (
	"reflect"
	"testing"
)

func TestCombinationOfInput(t *testing.T) {
	type args struct {
		in     []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{

			name: "Test",
			args: args{
				in:     []int{2, 3, 6, 7},
				target: 7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "[2,3,5]",
			args: args{
				in:     []int{2, 3, 5},
				target: 8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.in, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationOfInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
