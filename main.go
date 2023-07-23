package main

import (
	"encoding/binary"
	"fmt"
	"prefixed-distributed-ids/number_generator"
	"prefixed-distributed-ids/timestamp_generator"
)

func main() {
	idBytes, id, timestamp := BuildId()

    fmt.Printf("bytes are %x\n", idBytes)
    fmt.Printf("timestamp is %d\n", timestamp)
    fmt.Printf("id is %s\n", id)
}

func BuildId() ([]byte, string, int64) {
    var idBytes []byte
    tBuffer, timestamp := timestamp_generator.GenerateTimestampNumber()

    idBytes = append(idBytes, tBuffer[0], tBuffer[1], tBuffer[2], tBuffer[3])

    processBuffer, err := number_generator.GenerateBytes()

	if err != nil {
		panic(err)
	}

    idBytes = append(idBytes, processBuffer[0], processBuffer[1], processBuffer[2], processBuffer[3], processBuffer[4])

    counter, err := number_generator.InitializeCounter()
	counterBuffer := make([]byte, 2)
	binary.BigEndian.PutUint16(counterBuffer, uint16(counter))

    if err != nil {
		panic(err)
	}

    idBytes = append(idBytes, counterBuffer[0], counterBuffer[1])
    id := fmt.Sprintf("%x", idBytes)

    return idBytes, id, timestamp
}
