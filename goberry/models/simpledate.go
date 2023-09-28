package models

import "time"

type SimpleDate time.Time

const simpleDateLayout = "2006-01-02"

func (sd *SimpleDate) UnmarshalParam(param string) error {
	t, err := time.Parse(simpleDateLayout, param)
	if err != nil {
		return err
	}
	*sd = SimpleDate(t)
	return nil
}
