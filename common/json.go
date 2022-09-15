package common

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseBody(body io.Reader, value interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, value)
	if err != nil {
		return err
	}
	return nil
}
