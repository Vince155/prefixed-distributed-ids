package pregen

import (
	"testing"

    "github.com/stretchr/testify/assert"
)

func TestPrefixIdIsGenerated(t *testing.T) {
	pfix := "test"
	prefix, err := BuildId(pfix)

	assert.Equal(t, err, nil, "Error should be nil")
	assert.Greater(t, prefix.Ts, int64(0), "timestamp should be current")
	assert.NotEqual(t, prefix.ByteArr, []byte{}, "Bytes should not be empty")
	assert.Contains(t, prefix.Id, "test_", "Resulting ID should contain 'test_'")
}

func TestIDsAreUnique(t *testing.T) {
	pfix := "test"
	prefix1, _ := BuildId(pfix)
	prefix2, _ := BuildId(pfix)

	assert.NotEqual(t, prefix1.Id, prefix2.Id, "The strings should be unique")
}

func TestIDsContainGivenPrefix(t *testing.T) {
	pfix := "foo"
	prefix1, _ := BuildId(pfix)
	pfix = "bar"
	prefix2, _ := BuildId(pfix)

	assert.Contains(t, prefix1.Id, "foo", "First ID should contain 'foo'")
	assert.NotContains(t, prefix1.Id, "bar", "First ID should not contain 'bar'")
	assert.Contains(t, prefix2.Id, "bar", "Second ID should contain 'bar'")
	assert.NotContains(t, prefix2.Id, "foo", "Second ID should not contain 'foo'")
}

func TestPrefixMustBeShorterThan9Characters(t *testing.T) {
	pfix := "Superlongprefix"
	_, err := BuildId(pfix)

	assert.Equal(t, err.Error(), "please use a prefix shorter than 9 characters", "The two error message are equal")
}

func TestPrefixMustExist(t *testing.T) {
	pfix := ""
	_, err := BuildId(pfix)

	assert.Equal(t, err.Error(), "please enter a prefix", "The two error message are equal")
}

func TestNextIdCanBeGenerated(t *testing.T) {
	pfix := "Test"
	prefid, _ := BuildId(pfix)

	nextId, err := prefid.NextId(pfix)

	assert.NotEqual(t, nextId.ByteArr, prefid.ByteArr, "The two IDs should have unique bytes")
	assert.NotEqual(t, nextId.Id, prefid.Id, "The two IDs should have unique strings")
	assert.NotEqual(t, nextId.Ts, prefid.Ts, "The two IDs should have unique timestamps")
	assert.Equal(t, err, nil, "Error should be nil")
}
