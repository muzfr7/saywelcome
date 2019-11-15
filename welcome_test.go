package saywelcome

import "testing"

func TestWelcome(t *testing.T) {
	want := "Welcome.."
	got := Welcome()

	if want != got {
		t.Errorf("Welcome() = %q, want %q", got, want)
	}
}
