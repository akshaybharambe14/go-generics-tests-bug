package slices

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	type args[T any] struct {
		s []T
		i int
	}
	tests := []struct {
		name        string
		wantResult  []int
		args        args[int]
		wantDeleted bool
	}{
		{
			name:        "delete from empty slice",
			args:        args[int]{s: []int{}, i: 0},
			wantResult:  []int{},
			wantDeleted: false,
		},
		{
			name:        "delete out of bound index",
			args:        args[int]{s: []int{1, 2}, i: 5},
			wantResult:  []int{1, 2},
			wantDeleted: false,
		},
		{
			name:        "delete negative index",
			args:        args[int]{s: []int{1, 2}, i: -5},
			wantResult:  []int{1, 2},
			wantDeleted: false,
		},
		{
			name:        "delete valid index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 2},
			wantResult:  []int{1, 2, 1, 2, 3},
			wantDeleted: true,
		},
		{
			name:        "delete last index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 5},
			wantResult:  []int{1, 2, 3, 1, 2},
			wantDeleted: true,
		},
		{
			name:        "delete first index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 0},
			wantResult:  []int{2, 3, 1, 2, 3},
			wantDeleted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Delete(tt.args.s, tt.args.i)
			if !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("Delete() got = %v, want %v", got, tt.wantResult)
			}
			if got1 != tt.wantDeleted {
				t.Errorf("Delete() got1 = %v, want %v", got1, tt.wantDeleted)
			}
		})
	}
}
