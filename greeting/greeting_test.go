package greeting

import "testing"

func TestHello(t *testing.T) {
	want := "Halo, Go Developer! Salam dari private-greeting v0.1.0."
	if got := Hello("Go Developer"); got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}
