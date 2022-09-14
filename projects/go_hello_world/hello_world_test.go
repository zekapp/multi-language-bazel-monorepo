package go_hellow_world

import "testing"

func TestGreeter(t *testing.T)  {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("expected %q but not got %q", expected, actual)
	}
}