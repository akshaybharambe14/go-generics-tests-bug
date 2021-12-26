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
