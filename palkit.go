package palkit

import (
	"errors"
	"strconv"
)

func Atoi64(toConvert string) (int64, error) {
	return strconv.ParseInt(toConvert, 0, 64)
}

func JsonNumberToInt(toConvert interface{}) (int64, error) {
	switch toConvert.(type) {

	case float64:
		return int64(toConvert.(float64)), nil
		break

	default:
		return 0, errors.New("Wrong format, was expecting number")
		break
	}
	return 0, nil
}
