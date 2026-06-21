package Number

import "testing"

// TestCachedInt white-box-tests the small-integer cache directly: the in-band
// hit, the out-of-band miss, and the non-integer guard (the last is defensive —
// the public call sites only pass integer values when isInt is true, so it is
// not reachable through the API).
func TestCachedInt(t *testing.T) {
	if _, ok := cachedInt(5); !ok {
		t.Fatal("a small integer should be cached")
	}
	if _, ok := cachedInt(1000); ok {
		t.Fatal("an out-of-band integer should not be cached")
	}
	if _, ok := cachedInt(3.5); ok {
		t.Fatal("a non-integer value must not be served from the int cache")
	}
}
