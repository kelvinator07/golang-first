package about_test

import (
	"testing"

	"about"
)

func test_about(t *testing.T) {
	if about.Name() != "first last name" {
		t.Fatal("Wrong name")
	}
}
