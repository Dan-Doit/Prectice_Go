package dicmap

import (
	"errors"
	"fmt"
)

type Dicts map[string]string

var (
	AddError error = errors.New("key is already exists")
	DelError error = errors.New("Can't find your key")
	UpdError error = errors.New("Can't update, check key first")
)

func (d Dicts) Check(key string) bool {
	_, exists := d[key]
	fmt.Println(d)
	return exists
}

func (d Dicts) AddDicts(key, value string) (Dicts, error) {
	if d.Check(key) {
		return d, AddError
	} else {
		d[key] = value
		return d, nil
	}
}

func (d Dicts) DelDicts(key string) (Dicts, error) {
	if d.Check(key) {
		delete(d, key)
		return d, nil
	} else {
		return nil, DelError
	}
}

func (d Dicts) UpdDicts(key, value string) (Dicts, error) {
	if d.Check(key) {
		d[key] = value
		return d, nil
	} else {
		return nil, UpdError
	}
}
