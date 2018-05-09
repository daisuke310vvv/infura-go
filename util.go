package infura

import (
	"fmt"
)

func toHexString(n int) string {
	return fmt.Sprintf("%0x", n)
}
