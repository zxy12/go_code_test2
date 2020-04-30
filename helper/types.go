package helper

import (
	"errors"
	"strconv"
)

const (
	UINT_MIN uint = 0
	UINT_MAX uint = ^UINT_MIN
)

const (
	INT_MAX int = int(^uint(0) >> 1)
	INT_MIN int = ^INT_MAX
)

const (
	UINT64_MIN uint64 = 0
	UINT64_MAX uint64 = ^UINT64_MIN
)

const (
	INT64_MAX int64 = int64(^uint64(0) >> 1)
	INT64_MIN int64 = ^INT64_MAX
)

type ComType string

func (t *ComType) Load(v interface{}) error {
	var str string
	switch v.(type) {
	case string:
		str = v.(string)
	case int:
		str = strconv.Itoa(v.(int))
	case int64:
		str = strconv.FormatInt(v.(int64), 10)
	case uint:
		i := v.(uint)
		str = strconv.FormatInt(int64(i), 10)
	case uint64:
		str = strconv.FormatUint(v.(uint64), 10)
	default:
		return errors.New("unsupport type")
	}
	c := ComType(str)
	*t = c
	return nil
}

func (t ComType) String() (string, error) {
	str := string(t)
	return str, nil
}

func (t ComType) Int() (int, error) {
	str := string(t)
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, err
}

func (t ComType) Uint() (uint, error) {
	str := string(t)
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint(i), err
}

func (t ComType) Int64() (int64, error) {
	str := string(t)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, err
}

func (t ComType) Uint64() (uint64, error) {
	str := string(t)
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, err
}
