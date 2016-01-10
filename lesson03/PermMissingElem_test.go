package lesson03_test

import (
	"reflect"
	"testing"

	"github.com/spring1843/go-codility/lesson03"
)

func TestPermMissingElem(t *testing.T) {
	solutionArray := lesson03.PermMissingElem(5, []int{3, 4, 4, 6, 1, 4, 4})

	if reflect.DeepEqual(solutionArray, []int{3, 2, 2, 4, 2}) != true {
		t.Error("missmatch for first case, expected 3,2,4,2 , got %d", solutionArray)
	}
}
