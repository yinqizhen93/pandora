package utils

import (
	"encoding/json"
	"time"
)

type LocalTime time.Time

// 确保编译时，可以判断LocalTime是否实现json.Unmarshaler接口
var _ json.Unmarshaler = &LocalTime{}

var _ json.Marshaler = &LocalTime{}

func (lt *LocalTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*lt = LocalTime(t)
	return nil
}

func (lt *LocalTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*lt).Format("2006-01-02"))
}
