package geetest

import "testing"

func TestGeetest(t *testing.T) {
	g := New("id", "key", true)

	_, err := g.Register()
	if err != nil {
		t.Errorf("expect error nil but get %v", err)
	}
}
