package test

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{12, 35, 37},
	}

	for _, tt := range tests {
		if actural := calTriangle(tt.a, tt.b); actural != tt.c {
			t.Errorf("wrong:%d,%d,%d", actural, tt.a, tt.b)
		}
	}
}
