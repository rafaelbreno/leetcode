package calculator_test

import (
	"go-calc"
	"reflect"
	"testing"
)

func TestSimpleSum(t *testing.T) {
	t.Helper()

	want := 3

	if got := calculator.Parse("2+1"); !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}
}
func TestSimpleSubtraction(t *testing.T) {
	t.Helper()

	want := 3

	if got := calculator.Parse("5-2"); !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}
}
func TestSimpleMultiplication(t *testing.T) {
	t.Helper()

	want := 6

	if got := calculator.Parse("2*3"); !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}
}
func TestSimpleFraction(t *testing.T) {
	t.Helper()

	want := 3

	if got := calculator.Parse("6/2"); !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}
}
