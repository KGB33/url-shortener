package base62

import (
	"errors"
	"math"
	"strings"
)

const basis = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const dimension = uint64(len(basis))

func Encode(num uint64) string {
	var builder strings.Builder
	for ; num > 0; num = num / dimension {
		builder.WriteByte(basis[(num % dimension)])
	}
	return builder.String()
}

func Decode(encoded string) (uint64, error) {
	var num uint64

	for i, r := range encoded {
		basisPos := strings.IndexRune(basis, r)

		if basisPos == -1 {
			return num, errors.New("Invaid rune: " + string(r))
		}
		num += uint64(basisPos) * uint64(math.Pow(float64(dimension), float64(i)))
	}
	return num, nil
}
