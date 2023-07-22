package number_generator

import (
	"testing"

    "github.com/stretchr/testify/assert"
)

func TestNumberIsGenerated(t *testing.T) {
    value, err := generateBytes()

    assert.NotEqual(t, 0, value, "value should not be -1")
    assert.Equal(t, err, nil, "Error should be nil")
}

func TestNumbersGeneratedAreUnique(t *testing.T) {
    value1, err1 := generateBytes()
    value2, err2 := generateBytes()

    assert.NotEqual(t, value1, value2, "The two values should be different")
    assert.Equal(t, err1, err2, "Both errors should be nil")
}
