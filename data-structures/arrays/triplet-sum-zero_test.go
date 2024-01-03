package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripletSumZero(t *testing.T) {

	is := assert.New(t)

	tests := []struct {
		name   string
		want   [][]int
		expect [][]int
	}{
		{
			name:   "triplet not found",
			want:   TripletSumZero([]int{}),
			expect: [][]int{},
		},
		{
			name:   "find two triplets",
			want:   TripletSumZero([]int{-1, 0, 1, 2, -1, -4}),
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:   "find one triplet",
			want:   TripletSumZero([]int{0, 0, 0}),
			expect: [][]int{{0, 0, 0}},
		},
		{
			name:   "find five triplets",
			want:   TripletSumZero([]int{1, 2, 3, 0, -1, -2, -3}),
			expect: [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			is.Equal(tc.expect, tc.want)
		})
	}
}
