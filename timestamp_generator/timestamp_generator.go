package timestamp_generator

import (
	"encoding/binary"
	"time"
)

func GenerateTimestampNumber() ([]byte, int64) {
    now := time.Now()
    timestamp := now.Unix()
    buffer := make([]byte, 4)
    binary.BigEndian.PutUint32(buffer, uint32(timestamp))

    return buffer, timestamp
}