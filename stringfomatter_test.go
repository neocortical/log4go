package log5go

import (
	"testing"
)

func TestNewStringFormatter(t *testing.T) {
	// all the stuff
	sf := NewStringFormatter("%t %l/%L %p (%c:%n): %艾未未 %m %d %% junk %")

	expected := []string{"%t", " ", "%l", "/L ", "%p", " (", "%c", ":", "%n", "): 艾未未 ", "%m", " d ", "%%", " junk "}

	if !testEq(sf.parts, expected) {
		t.Errorf("expected \n%v, but got \n%v", expected, sf.parts)
	}

	// empty format string
	sf = NewStringFormatter("")

	expected = []string{}

	if !testEq(sf.parts, expected) {
		t.Errorf("expected \n%v, but got \n%v", expected, sf.parts)
	}
}

func TestStringFormatterParse(t *testing.T) {
	sf := NewStringFormatter("%t %l %p (%c:%n): %m %%艾未未")
	actual := sf.Format("2014", "INFO", "艾未未", "acme.go", 123, "hello?")
	expected := "2014 INFO 艾未未 (acme.go:123): hello? %艾未未"
	if expected != string(actual) {
		t.Errorf("expected %s but got %s", expected, actual)
	}

	sf = NewStringFormatter("")
	actual = sf.Format("2014", "INFO", "艾未未", "acme.go", 123, "hello?")
	expected = ""
	if expected != string(actual) {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func testEq(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}