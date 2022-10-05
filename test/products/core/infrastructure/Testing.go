package infrastructure

import "testing"

func Assert(t *testing.T, got any, want any) {
	if got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
