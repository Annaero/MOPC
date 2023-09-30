package models

import (
	"fmt"
	"strings"
	"time"
)

const DATE_LAYOUT = "2006-01-02"

type Date struct {
	time.Time
}

func (t Date) MarshalJSON() ([]byte, error) {
	date := t.Time.Format(DATE_LAYOUT)
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse(DATE_LAYOUT, s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
