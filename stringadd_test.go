package stringadd

import (
	"testing"
)

func TestAdd_EmptyString_Returns0(t *testing.T) {
	if Add("") != 0 {
		t.Error("failing")
	}
}
