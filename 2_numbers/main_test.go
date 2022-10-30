package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumbers(t *testing.T) {
	for _, tc := range []struct {
		input int
	}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
		{10},
		{11},
		{19},
		{20},
		{23},
		{25},
		{33},
		{57},
		{68},
		{100},
		{101},
		{123},
		{376},
		{1279},
		{98764},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			var wantascending, wantdescending []int
			for i, j := 1, tc.input; i <= tc.input; i, j = i+1, j-1 {
				wantascending, wantdescending = append(wantascending, i), append(wantdescending, j)
			}
			if gotascending, gotdescending := Numbers(tc.input); !reflect.DeepEqual(gotascending, wantascending) || !reflect.DeepEqual(gotdescending, wantdescending) {
				t.Errorf("Numbers(%v) = %v %v, want %v %v", tc.input, gotascending, gotdescending, wantascending, wantdescending)
			}
		})
	}
}
