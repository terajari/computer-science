package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquilibrium(t *testing.T) {

	is := assert.New(t)

	tests := []struct {
		name   string
		want   int
		expect int
	}{
		{
			name:   "Equilibrium not found",
			want:   Equilibrium([]int{1, 2, 3}),
			expect: -1,
		},
		{
			name:   "Equilibrium found",
			want:   Equilibrium([]int{-7, 1, 5, 2, -4, 3, 0}),
			expect: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			is.Equal(tc.expect, tc.want)
		})
	}
}
