package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"time"
)

type Carbon struct {
	carbon.Carbon
}

func (t Carbon) Value() (driver.Value, error) {
	if !t.IsValid() {
		return nil, nil
	}
	return t.StdTime(), nil
}

func (t *Carbon) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Carbon{carbon.CreateFromStdTime(value)}
		return nil
	}
	carbon.CreateFromStdTime(value)
	return fmt.Errorf("can not convert %v to timestamp", v)
}
