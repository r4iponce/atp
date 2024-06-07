package atp

import (
	"errors"
	"testing"
	"time"
)

func TestUtils(t *testing.T) {
	t.Parallel()

	testCheckCharNotRedundantTrue(t)
	testCheckCharNotRedundantFalse(t)
	testParseExpirationFull(t)
	testParseExpirationMissing(t)
	testParseExpirationWithCaps(t)
	testParseExpirationNull(t)
	testParseExpirationNegative(t)
	testParseExpirationInvalidRedundant(t)
}

func testCheckCharNotRedundantTrue(t *testing.T) { // Test checkCharRedundant with redundant char
	t.Helper()
	want := true
	got := checkCharRedundant("2d1h3md4h7s", "h")
	if got != want {
		t.Fatalf("Error in parseExpirationFull, want: %t, got: %t", want, got)
	}
}

func testCheckCharNotRedundantFalse(t *testing.T) { // Test checkCharRedundant with not redundant char
	t.Helper()
	want := false
	got := checkCharRedundant("2d1h3m47s", "h")
	if got != want {
		t.Fatalf("Error in parseExpirationFull, want: %t, got: %t", want, got)
	}
}

func testParseExpirationFull(t *testing.T) {
	t.Helper()
	got, _ := ParseDuration("2d1h3m47s")
	want := time.Duration(176627000000000)
	if want != got {
		t.Fatalf("Error in parseExpirationFull, want: %s got: %s", want, got)
	}
}

// testParseExpirationMissing test with missing unit.
func testParseExpirationMissing(t *testing.T) {
	t.Helper()
	got, _ := ParseDuration("1h47s")
	want := 3647 * time.Second
	if want != got {
		t.Fatalf("Error in parseExpirationFull, want: %s got: %s", want, got)
	}
}

// testParseExpirationWithCaps verify case-insensitive.
func testParseExpirationWithCaps(t *testing.T) {
	t.Helper()
	got, _ := ParseDuration("2D1h3M47s")
	want := 176627 * time.Second
	if got != want {
		t.Fatalf("Error in parseExpirationFull, want: %s got: %s", want, got)
	}
}

func testParseExpirationNull(t *testing.T) {
	t.Helper()
	got, _ := ParseDuration("0")
	want := time.Duration(0)
	if got != want {
		t.Fatalf("Error in ParseExpirationFull, want: %s got: %s", want, got)
	}
}

func testParseExpirationNegative(t *testing.T) {
	t.Helper()
	_, got := ParseDuration("-42h1m4s")
	if got == nil {
		t.Fatal("testParseExpirationNegative: ParseDuration is supposed to crash with negative value")
	}
}

// testParseExpirationInvalidRedundant invalid duration and redundant char.
func testParseExpirationInvalidRedundant(t *testing.T) {
	t.Helper()
	_, got := ParseDuration("8h42h1m1h4s")
	want := errRedundantChar
	if !errors.Is(got, errRedundantChar) {
		t.Fatalf("Error in ParseExpirationFull, want: %s got: %s", want, got)
	}
}
