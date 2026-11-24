package main

import "testing"

func TestSearch(t *testing.T) {
	root := tn(8, tn(3, tn(1, nil, nil), tn(6, tn(4, nil, nil), tn(7, nil, nil))), tn(10, nil, tn(14, tn(13, nil, nil), nil)))
	for _, v := range []int{8, 3, 1, 6, 4, 7, 10, 14, 13} {
		if got := Search(root, v); got == nil || got.Val != v {
			t.Errorf("Search(%d) failed", v)
		}
	}
	for _, v := range []int{0, 5, 9, 99} {
		if got := Search(root, v); got != nil {
			t.Errorf("Search(%d) = %v, want nil", v, got)
		}
	}
	if Search(nil, 1) != nil {
		t.Error("empty tree should return nil")
	}
}
