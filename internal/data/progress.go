package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidProgressFormat = errors.New("invalid progress format")

type Progress int32

func (p Progress) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d %%", p)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (p *Progress) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidProgressFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "%" {
		return ErrInvalidProgressFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidProgressFormat
	}

	*p = Progress(i)

	return nil
}
