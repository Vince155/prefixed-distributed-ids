package tsgen

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBytesAreUnique(t *testing.T) {
    byteBuffer1, timestamp1 := GenerateTimestampNumber()
    time.Sleep(time.Second)
    byteBuffer2, timestamp2 := GenerateTimestampNumber()

    assert.NotEqual(t, byteBuffer1, byteBuffer2, "The two byte slices should be different")
    assert.NotEqual(t, timestamp1, timestamp2, "The two timestamps should be different")
}
