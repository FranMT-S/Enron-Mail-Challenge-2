package shared

import (
	"fmt"
	"time"
)

func ParseDate(value string) (date time.Time, err error) {
	t, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700 (MST)", value)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}

	return t, nil
}
