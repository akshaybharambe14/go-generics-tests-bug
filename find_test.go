package slices

import "testing"

func TestFind(t *testing.T) {
	type args[T comparable] struct {
		v T
		s []T
	}

	type want struct {
		index int
		found bool
	}

	tests := []struct {
		name      string
		args      args[int]
		wantIndex int
		wantFound bool
	}{
		{
			name: "element exists in slice",
			args: args[int]{
				v: 1,
				s: []int{1, 2, 3, 4, 5},
			},
			wantIndex: 0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := Find(tt.args.s, tt.args.v)
			if gotIndex != tt.wantIndex {
				t.Errorf("Find() got = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Find() got1 = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
