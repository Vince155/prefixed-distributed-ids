package timestamp_generator

import (
    "time"
    "encoding/binary"
)

func generateTimestampNumber() []byte {
    now := time.Now()
    timestamp := now.Unix()
    buffer := make([]byte, 4)
    binary.BigEndian.PutUint64(buffer, uint64(timestamp))

    return buffer
}