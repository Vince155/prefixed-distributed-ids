package number_generator

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/klauspost/cpuid/v2"
)

type CPUInfo struct {
	BrandName string
	PhysicalCores uint8
	LogicalCores uint8
	CPUFamilyNumber int
	CPUModel int
	VendorId string
	Frequency int64
	L1DataCache int
	L1InstructionCache int
	L2Cache int
	L3Cache int
}

type Seed struct {
    Value string
    CpuInfo CPUInfo
}

const randNumChars = "1234567890"

func generateRandomNumbers(idLength int) (string, error) {
	buffer := make([]byte, idLength)
    _, err := rand.Read(buffer)

    if err != nil {
        return "", err
    }

    charsLength := len(randNumChars)

    for i := 0; i < len(buffer); i++ {
        buffer[i] = randNumChars[int(buffer[i]) % charsLength]
    }

    return string(buffer), nil
}

func generateBytes() ([]byte, error) {
    idLength := 6
	randomValue, err := generateRandomNumbers(idLength)

    if err != nil {
        return []byte{}, err
    }
	
	cpuInfo := CPUInfo{
		BrandName: cpuid.CPU.BrandName,
		PhysicalCores: uint8(cpuid.CPU.PhysicalCores),
		LogicalCores: uint8(cpuid.CPU.LogicalCores),
		CPUFamilyNumber: cpuid.CPU.Family,
		CPUModel: cpuid.CPU.Model,
		VendorId: cpuid.CPU.VendorID.String(),
        Frequency: cpuid.CPU.Hz,
        L1DataCache: cpuid.CPU.Cache.L1D,
        L1InstructionCache: cpuid.CPU.Cache.L1I,
        L2Cache: cpuid.CPU.Cache.L2,
        L3Cache: cpuid.CPU.Cache.L3,
	}
    seed := Seed{
        Value: randomValue,
        CpuInfo: cpuInfo,
    }
    seedBytes, err := json.Marshal(seed)

    if err != nil {
       fmt.Println("error: ", err)

       return []byte{}, err
    }

    h := sha256.New()
    h.Write(seedBytes)
    hashedBytes := h.Sum(nil)

	return hashedBytes[:5], nil
}
