package rajirec

import (
	"testing"
)

func TestMap1(t *testing.T) {
	m := NewMap()
	m.Set("ab", "a", "b")
	output := m.Get("a", "b")
	if val, ok := output.GetValue().(string); !ok {
		t.Errorf("Value %v is not convertible to string", val)
	} else if val != "ab" {
		t.Errorf("output: %s, expected: %s", val, "ab")
	}
}

func TestMap2(t *testing.T) {
	m := NewMap()
	output := m.Get("a", "b")
	if val := output.GetValue(); val != nil {
		t.Errorf("output: %s, expected: nil", val)
	}
}

