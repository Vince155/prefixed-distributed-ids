package tsgen

import (
	"encoding/binary"
	"time"
)

func GenerateTimestampNumber() ([]byte, int64) {
    now := time.Now().UTC()
    timestamp := now.UnixMilli()
    buffer := make([]byte, 4)
    binary.BigEndian.PutUint32(buffer, uint32(timestamp))

    return buffer, timestamp
}