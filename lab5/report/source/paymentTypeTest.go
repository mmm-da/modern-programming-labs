package payment

import "testing"

func TestPaymentTypeString(t *testing.T) {
	expected := "Regular"
	received := Regular.String()
	if received != expected {
		t.Fatalf("expected %s, received %s", expected, received)
	}
}
