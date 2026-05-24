package hello 

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, test!"
	got := Say("test w")

	if want != got {
		t.Errorf("Wanted %s, got %s",want , got)
	}
}