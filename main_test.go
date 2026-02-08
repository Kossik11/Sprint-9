package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantNil bool
		wantLen int
	}{
		{
			name:    "size zero",
			size:    0,
			wantNil: true,
			wantLen: 0,
		},
		{
			name:    "negative size",
			size:    -10,
			wantNil: true,
			wantLen: 0,
		},
		{
			name:    "positive size",
			size:    1000,
			wantNil: false,
			wantLen: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if tt.wantNil {
				require.Nil(t, data)
				return
			}

			require.NotNil(t, data)
			require.Len(t, data, tt.wantLen)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "nil slice", data: nil, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "many elements", data: []int{3, 1, 9, 2, 9, 5}, want: 9},
		{name: "all equal", data: []int{4, 4, 4}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "nil slice", data: nil, want: 0},
		{name: "small slice", data: []int{1, 5, 2, 4, 3}, want: 5},
		{name: "many elements", data: []int{3, 1, 9, 2, 9, 5}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			require.Equal(t, tt.want, got)
		})
	}
}
