package int2bin

import (
	"fmt"
)

func Bin(n interface{}) (string, error) {
	switch v := n.(type) {
	case int:
		return binSigned(int64(v), 32<<(^uint(0)>>32&1)), nil
	case int8:
		return binSigned(int64(v), 8), nil
	case int16:
		return binSigned(int64(v), 16), nil
	case int32:
		return binSigned(int64(v), 32), nil
	case int64:
		return binSigned(v, 64), nil
	case uint:
		return binUnsigned(uint64(v), 32<<(^uint(0)>>32&1)), nil
	case uint8:
		return binUnsigned(uint64(v), 8), nil
	case uint16:
		return binUnsigned(uint64(v), 16), nil
	case uint32:
		return binUnsigned(uint64(v), 32), nil
	case uint64:
		return binUnsigned(v, 64), nil
	}

	return "", fmt.Errorf("unsupported type %T", n)
}

func binSigned(n int64, len uint) string {
	var activeBit int64
	var cnt = 1
	var binString [80]byte
	i := 0

	// sign bit
	binString[i] = byte((n>>63)&1) + '0'

	for activeBit = 1 << (len - 2); activeBit > 0; activeBit /= 2 {
		i++

		if (n & activeBit) == activeBit {
			binString[i] = '1'
		} else {
			binString[i] = '0'
		}

		cnt++
		if cnt == 4 {
			i++
			binString[i] = ' '
			cnt = 0
		}
	}

	return string(binString[0:i])
}

func binUnsigned(n uint64, len uint) string {
	var activeBit uint64
	var cnt = 0
	var binString [80]byte
	i := 0

	for activeBit = 1 << (len - 1); activeBit > 0; activeBit /= 2 {
		if (n & activeBit) == activeBit {
			binString[i] = '1'
		} else {
			binString[i] = '0'
		}

		cnt++
		if cnt == 4 {
			i++
			binString[i] = ' '
			cnt = 0
		}

		i++
	}

	return string(binString[0 : i-1])
}
