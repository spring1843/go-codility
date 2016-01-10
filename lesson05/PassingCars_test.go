package lesson05_test

import (
	"testing"

	"github.com/spring1843/go-codility/lesson05"
)

func TestPassingCars(t *testing.T) {
	actual := lesson05.PassingCars([]int{0, 1, 0, 1, 1})
	expected := 5
	if actual != expected {
		t.Errorf("PermMissingElm failed, expected %d, got %d", expected, actual)
	}

	actual = lesson05.PassingCars([]int{0, 1, 0, 1, 1, 1})
	expected = 7
	if actual != expected {
		t.Errorf("PermMissingElm failed, expected %d, got %d", expected, actual)
	}
}
