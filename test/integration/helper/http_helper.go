package helper

import (
	"bytes"
	"encoding/json"
	"io"
)

func WriteBody(body []byte) *bytes.Reader {
	return bytes.NewReader(body)
}

func ParseRequest[T any](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}

	return nil
}
