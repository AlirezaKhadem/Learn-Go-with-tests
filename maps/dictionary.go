package maps

import "errors"

type Dictionary map[string]string

var (
	KeyDoesNotExistError  = errors.New("key does not exist")
	KeyAlreadyExistsError = errors.New("key already exists")
	NotFoundError         = errors.New("could not find the key you were looking for")
)

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", KeyDoesNotExistError
	}

	return value, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, ok := d[key]
	if ok {
		return KeyAlreadyExistsError
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, newValue string) error {
	_, ok := d[key]
	if !ok {
		return NotFoundError
	}

	d[key] = newValue
	return nil
}
