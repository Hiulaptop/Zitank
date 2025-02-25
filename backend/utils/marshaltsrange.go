package utils

import (
	"fmt"

	"github.com/jackc/pgtype"
)

type tsr pgtype.Tsrange

func (t *tsr) UnmarshalJSON(data []byte) error {
	fmt.Println(data)
	t.Lower.Scan("2025-02-20 00:00:00")
	t.Upper.Scan("2025-02-27 23:59:59")
	return nil
}
