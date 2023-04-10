package maps

import "errors"

type Dictionary map[string]string

var KeyNotExistError = errors.New("key not exist")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", KeyNotExistError
	}

	return value, nil
}
