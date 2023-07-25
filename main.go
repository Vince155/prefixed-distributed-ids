package pregen

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/Vince155/prefixed-distributed-ids/numgen"
	"github.com/Vince155/prefixed-distributed-ids/tsgen"
)

const maxCounter uint = 999

type PrefId struct {
    ByteArr []byte
    Id string
    Ts int64
}

func BuildId(prefix string) (*PrefId, error) {
    if len(prefix) > 8 {
        return nil, errors.New("please use a prefix shorter than 9 characters")
    }

    if len(prefix) == 0 {
        return nil, errors.New("please enter a prefix")
    }

    var idBytes []byte
    tBuffer, timestamp := tsgen.GenerateTimestampNumber()

    idBytes = append(idBytes, tBuffer[0], tBuffer[1], tBuffer[2], tBuffer[3])

    processBuffer, err := numgen.GenerateBytes()

	if err != nil {
		panic(err)
	}

    idBytes = append(idBytes, processBuffer[0], processBuffer[1], processBuffer[2], processBuffer[3], processBuffer[4])

    counter, err := numgen.InitializeCounter()
	counterBuffer := make([]byte, 2)
	binary.BigEndian.PutUint16(counterBuffer, uint16(counter))

    if err != nil {
		panic(err)
	}

    idBytes = append(idBytes, counterBuffer[0], counterBuffer[1])
    strid := fmt.Sprintf("%x", idBytes)
    prefixedId := prefix + "_" + strid

    prefid := &PrefId{idBytes, prefixedId, timestamp}

    return prefid, nil
}
