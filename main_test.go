package pregen

import (
	"testing"

    "github.com/stretchr/testify/assert"
)

func TestPrefixIdIsGenerated(t *testing.T) {
	pfix := "test"
	b, str, ts, err := BuildId(pfix)

	assert.Equal(t, err, nil, "Error should be nil")
	assert.Greater(t, ts, int64(0), "timestamp should be current")
	assert.NotEqual(t, b, []byte{}, "Bytes should not be empty")
	assert.Contains(t, str, "test_", "Resulting ID should contain 'test_'")
}

func TestIDsAreUnique(t *testing.T) {
	pfix := "test"
	_, str1, _, _ := BuildId(pfix)
	_, str2, _, _ := BuildId(pfix)

	assert.NotEqual(t, str1, str2, "The strings should be unique")
}

func TestIDsContainGivenPrefix(t *testing.T) {
	pfix := "foo"
	_, str1, _, _ := BuildId(pfix)
	pfix = "bar"
	_, str2, _, _ := BuildId(pfix)

	assert.Contains(t, str1, "foo", "First ID should contain 'foo'")
	assert.NotContains(t, str1, "bar", "First ID should not contain 'bar'")
	assert.Contains(t, str2, "bar", "Second ID should contain 'bar'")
	assert.NotContains(t, str2, "foo", "Second ID should not contain 'foo'")
}

func TestPrefixMustBeShorterThan9Characters(t *testing.T) {
	pfix := "Superlongprefix"
	_, _, _, err := BuildId(pfix)

	assert.Equal(t, err.Error(), "please use a prefix shorter than 9 characters", "The two error message are equal")
}

func TestPrefixMustExist(t *testing.T) {
	pfix := ""
	_, _, _, err := BuildId(pfix)

	assert.Equal(t, err.Error(), "please enter a prefix", "The two error message are equal")
}
