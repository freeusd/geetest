package geetest

import "testing"

func TestGeetest(t *testing.T) {
	g := New("id", "key", true)

	if actual, expected := g.registrationURL(), "https://api.geetest.com/register.php?gt=id"; actual != expected {
		t.Errorf("expected registration url %s but get %s", expected, actual)
	}

	if actual, expected := g.validationURL(), "https://api.geetest.com/validate.php"; actual != expected {
		t.Errorf("expected validation url %s but get %s", expected, actual)
	}

	_, err := g.Register()
	if err != nil {
		t.Errorf("expect error nil but get %v", err)
	}
}
