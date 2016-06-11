package geetest

import (
	"testing"
	"time"
)

func TestGeetest(t *testing.T) {
	g := New("id", "key", true, time.Second, time.Second, 8)

	_, err := g.Register()
	if err != nil {
		t.Errorf("expect error nil but get %v", err)
	}
}
