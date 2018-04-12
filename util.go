package infura

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func EncodeToHex(s []byte) string {
	dst := make([]byte, hex.EncodedLen(len(s)))
	hex.Encode(dst, s)
	return string(dst)
}

func DecodeToInt(hexString string) (int64, error) {
	s := Remove0x(hexString)
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func ToHexString(n int) string {
	return fmt.Sprintf("%0x", n)
}

func Remove0x(s string) string {
	return strings.Replace(s, "0x", "", 1)
}
