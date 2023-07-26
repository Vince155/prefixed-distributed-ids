package pregen

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/Vince155/prefixed-distributed-ids/numgen"
	"github.com/Vince155/prefixed-distributed-ids/tsgen"
)

const maxCounter uint = 999

var counter int
var lts int64

type PrefId struct {
    ByteArr []byte
    Id string
    Ts int64
}

func (prefid *PrefId) NextId(prefix string) (*PrefId, error) {
    if len(prefix) > 8 {
        return nil, errors.New("please use a prefix shorter than 9 characters")
    }

    if len(prefix) == 0 {
        return nil, errors.New("please enter a prefix")
    }

    counter++
    counter = counter & int(maxCounter)
    cterBuffer := make([]byte, 2)
	binary.BigEndian.PutUint16(cterBuffer, uint16(counter))
    tBuffer, cts := tsgen.GenerateTimestampNumber()


    if cts < lts {
        return nil, errors.New("invalid timestamp")
    }

    if cts == lts {
        tBuffer, cts = waitNextMill()
    }

    lts = cts

    idBytes := prefid.ByteArr

    idBytes[0] = tBuffer[0]
    idBytes[1] = tBuffer[1]
    idBytes[2] = tBuffer[2]
    idBytes[3] = tBuffer[3]
    idBytes[9] = cterBuffer[0]
    idBytes[10] = cterBuffer[1]

    strid := fmt.Sprintf("%x", idBytes)
    prefixedId := prefix + "_" + strid

    prefid.ByteArr = idBytes
    prefid.Id = prefixedId
    prefid.Ts = cts

    return prefid, nil
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

    cter, err := numgen.InitializeCounter()

    if err != nil {
		panic(err)
	}

	counterBuffer := make([]byte, 2)
	binary.BigEndian.PutUint16(counterBuffer, uint16(cter))

    counter = cter

    idBytes = append(idBytes, counterBuffer[0], counterBuffer[1])
    strid := fmt.Sprintf("%x", idBytes)
    prefixedId := prefix + "_" + strid

    prefid := &PrefId{idBytes, prefixedId, timestamp}

    return prefid, nil
}

func waitNextMill() ([]byte, int64) {
    tBuffer, cts := tsgen.GenerateTimestampNumber()

    for cts == lts {
        tBuffer, cts = tsgen.GenerateTimestampNumber()
    }

    return tBuffer, cts
}
