package main

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/Vince155/prefixed-distributed-ids/number_generator"
	"github.com/Vince155/prefixed-distributed-ids/timestamp_generator"
)

func main() {
    b, s, t, err := BuildId("test")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("str %s\n", s)
    fmt.Printf("byte %s\n", b)
    fmt.Printf("ts %d\n", t)
}

func BuildId(prefix string) ([]byte, string, int64, error) {
    if len(prefix) > 8 {
        return []byte{}, "", -1, errors.New("please use a prefix shorter than 9 characters")
    }

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
    prefixedId := prefix + "_" + id

    return idBytes, prefixedId, timestamp, nil
}
