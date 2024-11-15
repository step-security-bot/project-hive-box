package handlers

import "testing"

func TestPrintHello(t *testing.T) {
	msg := "hello"

	if msg != PrintHello() {
		t.Fatalf("got %s want 'hello'", PrintHello())
	}
}
