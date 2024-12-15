package diff_test

import (
	diff "my_test/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	result := diff.Diff(5, 4)
	expected := 1
	assert.Equal(t, expected, result, "Diff(5, 4) should equals 1")
}
