package Number

import "testing"

// TestIntern white-box-tests the small-integer cache directly: the in-band hit
// returns the shared instance, and the out-of-band and non-integer paths return
// a freshly built Number. Every path must yield a non-nil Interface — intern
// must never return a nil Number, so that the Null-Object contract holds even
// on a cache miss (the branch the nonnil analyzer guards).
func TestIntern(t *testing.T) {
	// In-band integer: the shared, interned instance.
	hit := intern(5, true)
	if hit == nil {
		t.Fatal("intern must never return nil")
	}
	if hit != smallInts[5-cacheMin] {
		t.Fatal("an in-band small integer should return the shared instance")
	}
	if !hit.IsInt() || hit.IsNull() {
		t.Fatal("an in-band small integer should be a non-null integer")
	}

	// Out-of-band integer: a fresh, non-cached, non-nil instance.
	miss := intern(1000, true)
	if miss == nil {
		t.Fatal("an out-of-band integer must still return a non-nil Number")
	}
	if !miss.IsInt() {
		t.Fatal("an out-of-band integer should still report IsInt() == true")
	}
	for i := range smallInts {
		if miss == smallInts[i] {
			t.Fatal("an out-of-band integer must not alias a cached instance")
		}
	}

	// Non-integer value: a fresh, non-nil float Number (never served from the
	// int cache).
	fl := intern(3.5, false)
	if fl == nil {
		t.Fatal("a float value must still return a non-nil Number")
	}
	if !fl.IsFloat() {
		t.Fatal("a non-integer value must be a float Number, not an int")
	}
}
