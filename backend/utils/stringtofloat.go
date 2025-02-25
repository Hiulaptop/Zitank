package utils

import (
	"strconv"
	"strings"
)

type StringToFloat float64

func (foe *StringToFloat) UnmarshalJSON(data []byte) error {
	if string(data) == "\"\"" {
		if foe != nil {
			*foe = 0
		}
		return nil
	}
	num := strings.ReplaceAll(string(data), "\"", "")
	n, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return err
	}
	*foe = StringToFloat(n)
	return nil
}
