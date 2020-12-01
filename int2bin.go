package int2bin

import (
	"fmt"
)

func Bin(n interface{}) (string, error) {
	switch v := n.(type) {
	case int:
		return toBin(uint64(v), 32<<(^uint(0)>>32&1)), nil
	case int8:
		return toBin(uint64(v), 8), nil
	case int16:
		return toBin(uint64(v), 16), nil
	case int32:
		return toBin(uint64(v), 32), nil
	case int64:
		return toBin(uint64(v), 64), nil
	case uint:
		return toBin(uint64(v), 32<<(^uint(0)>>32&1)), nil
	case uint8:
		return toBin(uint64(v), 8), nil
	case uint16:
		return toBin(uint64(v), 16), nil
	case uint32:
		return toBin(uint64(v), 32), nil
	case uint64:
		return toBin(v, 64), nil
	}

	return "", fmt.Errorf("unsupported type %T", n)
}

func toBin(n uint64, len uint) string {
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
